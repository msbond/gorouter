APP=gorouter
APP_EXECUTABLE="$(APP)"

setup:
	glide install
	go build -o $(APP_EXECUTABLE)

compile:
	glide install
