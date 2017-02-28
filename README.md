# PersonalBackup
Personal Backup utility that zips the directories you specify and uploads it into Minio Object Server.

#Requirements
- Minio Object Server (Downloadable from [here](https://www.minio.io/downloads))
- Golang Development Environment to build the binary if you wish

#Configuration Details
A `backup-config.txt` is what the program reads to know which directories that it needs to zip along with Minio Server connection details. A sample `backup-config.txt` is provided with the project, you should modify it accordingly and place it in the same folder as the main go program. 

A sample `backup-config.txt` file is shown below:
~~~~
MinioHost = "localhost:9000"
BackupDirectories = [ "d:\\data1", "d:\\data2","d:\\data3" ]
BackupFileNamePrefixes = ["data1","data2","data3"]
AccessKey = "MINIO_ACCESS_KEY"
SecretKey = "MINIO_SECRET_KEY"
MinioBackupBucketName = "backups"
UseSSL = false`	
~~~~

You can specify as many directories that you would like to backup in the `BackupDirectories` property. The `BackupFileNamePrefixes` is the prefix attached to the Backup ZIP file that is generated, so that you can identify the backups as needed. Use a prefix of your choice. The number of entries that you specify for `BackupDirectories` and `BackupFileNamePrefixes` should be the same. Please make sure that you do so, there are no checks at this point in the program.

The backup ZIP files will be upload to the MINIO Server and they shall be placed in the bucket name that you specify in the `MinioBackupBucketName` property. **This bucket needs to be created prior to running the application**.

#To Run the Program
- Ensure that you have updated the `backup-config.txt` file as per your requirements. Ensure that Minio Server is running and the bucket name that you specify is present in Minio Server.
- Import the Golang libraries that are required

  `go get -u github.com/minio/minio-go`
  
  `go get -u github.com/pierrre/archivefile/zip`
  
  `go get -u github.com/twinj/uuid`
- Run the program via the following command:
  
  `go run minio-backup.go`
  
If all goes well, you should find the ZIP files of your backup present in the Minio Object Server. Check it via the UI available at `http://MINIO_HOST:9000`  
