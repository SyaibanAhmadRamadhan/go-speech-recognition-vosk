# Golang Speech Recognition with Vosk and PortAudio

## Project Description
This project demonstrates how to implement real-time speech-to-text transcription in Golang using the Vosk speech recognition toolkit and the PortAudio library for capturing audio input. Vosk is an open-source speech recognition toolkit that provides excellent accuracy for various languages and works offline, making it a powerful tool for real-time transcription applications.

## Setup Instructions
To run this project, you need to set up the necessary dependencies on your system. The primary dependencies are the Vosk API and PortAudio library. Below are the setup instructions for different operating systems.

### Prerequisites
1. Golang: Ensure you have Golang installed on your system.
2. Vosk Model: Download a Vosk model for the language you want to recognize from [Vosk's official models page](https://alphacephei.com/vosk/models). Extract the model to a directory of your choice.

### Installing Dependencies
#### linux debian
```shell
sudo apt-get install portaudio19-dev pkg-config
```

#### macos
```shell
brew install portaudio
```

### install package
```shell
go mod downlaod
```

## Running the Project
### clone project
```shell
https://github.com/SyaibanAhmadRamadhan/go-speech-recognition-vosk.git
```

### and run 
#### run windows
```shell
make run-windows
```
#### run linux
```shell
make run-linux
```