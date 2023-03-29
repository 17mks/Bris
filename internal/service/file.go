package service

import (
	"fmt"
	"followup/internal/biz"
	http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/minio/minio-go/v6"
	"gogs.buffalo-robot.com/gogs/checksum"
	"gogs.buffalo-robot.com/gogs/module/lib"
	"gogs.buffalo-robot.com/gogs/module/models"
	mmodels "gogs.buffalo-robot.com/services/basis/models"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

type FileSaver interface {
	SaveFile(file io.Reader, fileName string) (code string, err error)
}

type ossSaver struct {
	client     *minio.Client
	bucketName string
	tmpPath    string
}

func (s *ossSaver) SaveFile(file io.Reader, fileName string) (code string, err error) {
	ext := filepath.Ext(fileName)
	filenameall := path.Base(fileName)
	filesuffix := path.Ext(fileName)
	fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
	newFileName := fmt.Sprintf("%s-%d%s", fileprefix, models.GetSFID(), ext)
	filePath := filepath.Join(s.tmpPath, fileName)
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return "", err
	}
	contentType, err := checksum.GetFileContentType(filePath)
	if err != nil {
		return "", err
	}
	defer os.Remove(filePath)

	_, err = s.client.FPutObject(s.bucketName, newFileName, filePath, minio.PutObjectOptions{
		//ContentType: "application/" + fileType,
		ContentType: contentType,
	},
	)
	code = fmt.Sprintf("%s/%s", s.bucketName, newFileName)

	return
}

type osFileSystemSaver struct {
	basePath string
}

func (s *osFileSystemSaver) SaveFile(file io.Reader, fileName string) (code string, err error) {
	distPath := filepath.Join(s.basePath, fileName)
	dist, err := os.Create(distPath)
	if err != nil {
		return "", err
	}
	defer dist.Close()

	_, err = io.Copy(dist, file)
	return
}

func NewFileSaver(data biz.DataRepo) FileSaver {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	tmpPath := filepath.Join(wd, "tmp")
	if !lib.Exists(filepath.Join(wd, "tmp")) {
		err := os.Mkdir(tmpPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &ossSaver{
		client:     data.GetMinion(),
		bucketName: mmodels.BaseBucket,
		tmpPath:    tmpPath,
	}
}

// 文件上传
func (s *FileService) UploadFile(ctx http.Context) error {
	req := ctx.Request()

	var response struct {
		Status    string      `json:"status"`
		RequestID string      `json:"requestID"`
		Data      interface{} `json:"confData"`
		Message   string      `json:"message"`
	}

	fileName := req.FormValue("name")
	if fileName == "" {
		response.Status = "fail"
		response.Message = "文件名为空"
		ctx.JSON(200, &response)
		return nil
	}
	file, _, err := req.FormFile("file")
	if err != nil {
		//s.log.Log(zapcore.ErrorLevel, err.Error()
		response.Status = "fail"
		response.Message = err.Error()
		ctx.JSON(200, &response)
		return nil
	}
	defer file.Close()

	newFileName, err := s.saver.SaveFile(file, fileName)

	if err != nil {
		response.Status = "fail"
		response.Message = err.Error()
	} else {
		response.Status = "success"
		response.Data = newFileName
	}

	return ctx.JSON(200, &response)
}

//文件下载

// var str = "http://192.168.100.132:9090/bfr/肺围手术期症状量表(PSA-Lung)-1638080442039107584.zip"

func OSSFileDownLoader() error {
	client, err := minio.New("192.168.100.132:9090", "root", "bfr123123", false)
	if err != nil {
		panic(err)
	}
	file, err := client.GetObject("bfr", "肺围手术期症状量表(PSA-Lung)-1638080442039107584.zip", minio.GetObjectOptions{})

	if err != nil {
		return err
	}
	localFile, err := os.Create("F:\\Kratos\\followup\\files\\1.zip")

	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := io.CopyN(localFile, file, stat.Size); err != nil {
		log.Fatalln(err)
	}

	err = lib.Unzip("F:\\Kratos\\followup\\files\\1.zip", "F:\\Kratos\\followup\\files")
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
