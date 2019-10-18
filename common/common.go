package common

import "path"

var (
	WORK_SPACE_ROOT string = "/root/my_docker_workspace"
	CONTAINER_FILE_SYSTEM_MOUNT_ROOT string = path.Join(WORK_SPACE_ROOT, "mnt")
	IMAGE_REGISTRY string = path.Join(WORK_SPACE_ROOT, "image_registry")
	STATUS_RUNING string = "running"
	STATUS_STOP string = "stopped"
	InformationFileName string = "information.json"
	LogFileName string = "log.log"
	ErrorLogFileName string = "error.log"
	DefaultContainerInformationLocation string = "/var/run/mydocker/"
	ENV_CONTAINER_PID string = "ENV_CONTAINER_PID"
	ENV_CONTAINER_EXEC_COMMAND string = "ENV_CONTAINER_EXEC_COMMAND"
	EXEC_PARENT_PROCESS_ID string = "EXEC_PARENT_PROCESS_ID"
	CONTAINER_FILE_SYSTEM_ROOT string = "/root/my_docker_workspace/containers"
	BUFFER_SIZE int = 1024

	NETWORK_INFORMATION_DIRECTORY = path.Join(DefaultContainerInformationLocation, "network")
	IPAM_ALLOCAT_SUBNET_DUMP_PATH = path.Join(NETWORK_INFORMATION_DIRECTORY, "ipam.json")
)

