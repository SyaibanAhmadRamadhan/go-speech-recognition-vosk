.PHONY: run-linux \
		run-windows

VOSK_BINARIES_LINUX_DIR=vosk-linux-x86_64-0.3.45
VOSK_MODEL_DIR=vosk-model-small-en-us-0.15

run-windows:
    # Check and download Vosk binaries if not present
	if [ ! -d "$(VOSK_BINARIES_LINUX_DIR)" ]; then \
		wget https://github.com/alphacep/vosk-api/releases/download/v0.3.45/vosk-linux-x86_64-0.3.45.zip; \
		unzip vosk-linux-x86_64-0.3.45.zip; \
		cp vosk-linux-x86_64-0.3.45/*.dll .
		cp vosk-linux-x86_64-0.3.45/*.h .
	fi
	# Check and download Vosk model if not present
	if [ ! -d "$(VOSK_MODEL_DIR)" ]; then \
		wget https://alphacephei.com/vosk/models/vosk-model-small-en-us-0.15.zip; \
		unzip vosk-model-small-en-us-0.15.zip; \
	fi
	VOSK_PATH=`pwd` LD_LIBRARY_PATH=$VOSK_PATH CGO_CPPFLAGS="-I $VOSK_PATH" CGO_LDFLAGS="-L $VOSK_PATH -lvosk -lpthread -dl" go run .

run-linux:
    # Check and download Vosk binaries if not present
	if [ ! -d "$(VOSK_BINARIES_LINUX_DIR)" ]; then \
		wget https://github.com/alphacep/vosk-api/releases/download/v0.3.45/vosk-linux-x86_64-0.3.45.zip; \
		unzip vosk-linux-x86_64-0.3.45.zip; \
	fi
	# Check and download Vosk model if not present
	if [ ! -d "$(VOSK_MODEL_DIR)" ]; then \
		wget https://alphacephei.com/vosk/models/vosk-model-small-en-us-0.15.zip; \
		unzip vosk-model-small-en-us-0.15.zip; \
	fi
	VOSK_PATH=`pwd`/$(VOSK_BINARIES_LINUX_DIR) LD_LIBRARY_PATH=$$VOSK_PATH CGO_CPPFLAGS="-I $$VOSK_PATH" CGO_LDFLAGS="-L $$VOSK_PATH" go run .