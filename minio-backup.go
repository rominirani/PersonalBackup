package main

import (
        "fmt"
        toml "github.com/BurntSushi/toml"
        minio "github.com/minio/minio-go"
        pzip "github.com/pierrre/archivefile/zip"
        uuid "github.com/twinj/uuid"
        "io/ioutil"
        "log"
        "os"
        "path/filepath"
        "strings"
        "time"
)

type Config struct {
        MinioHost              string
        BackupDirectories      []string
        BackupFileNamePrefixes []string
        AccessKey              string
        SecretKey              string
        MinioBackupBucketName  string
        UseSSL                 bool
}

func main() {

        //Read the file first
        data, err := ioutil.ReadFile("backup-config.txt")
        if err != nil {
                log.Fatalf("Error while reading the config file. Error %v", err)
        }

        var conf Config
        if _, err := toml.Decode(string(data), &conf); err != nil {
                log.Fatalf("Error while decoding the config file. %v", err)
        }

        // Initialize minio client object.
        minioClient, err := minio.New(conf.MinioHost, conf.AccessKey, conf.SecretKey, conf.UseSSL)
        if err != nil {
                log.Fatalln(err)
        }

        //Iterate through each of the directories to backup
        for idx, backupdirname := range conf.BackupDirectories {
                filename := strings.Join([]string{conf.BackupFileNamePrefixes[idx], time.Now().Format("2006-01-02T15-04-05T1-04-05"), uuid.NewV4().String(), "zip"}, ".")

                tmpDir, err := ioutil.TempDir("", "zip")
                if err != nil {
                        panic(err)
                }

                defer func() {
                        _ = os.RemoveAll(tmpDir)
                }()

                outFilePath := filepath.Join(tmpDir, filename)

                err = pzip.ArchiveFile(backupdirname, outFilePath, nil)
                if err != nil {
                        panic(err)
                }

                contentType := "application/zip"

                // Upload the zip file with FPutObject
                n, err := minioClient.FPutObject(conf.MinioBackupBucketName, filename, outFilePath, contentType)
                if err != nil {
                        log.Fatalln(err)
                }

                log.Printf("Successfully uploaded %s of size %d\n", filename, n)

        }
}