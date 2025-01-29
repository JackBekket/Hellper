**stt**
================

**Summary**
---------

The `stt` package is a Telegram bot API handler for voice messages. It provides functionality to retrieve and download voice message files, and then use them for speech-to-text recognition.

**Configuration**
----------------

* Environment variables:
	+ `AI_ENDPOINT`
	+ `VOICE_RECOGNITION_SUFFIX`
	+ `VOICE_RECOGNITION_MODEL`
* No flags or cmdline arguments are provided for configuration.

**Launch Options**
-----------------

The package can be launched as a Telegram bot API handler. To do this, you can use the `HandleVoiceMessage` function, which is the main entry point of the package.

**Edge Cases**
-------------

* The package does not provide any specific edge cases for launching. However, it is assumed that the package will be used as a Telegram bot API handler, and the `HandleVoiceMessage` function will be called with a Telegram message object and a bot API object.

**File Structure**
----------------

* `stt.go`

**Code Relations**
-----------------

The code is structured in a way that the `GetFileURL` function retrieves the file URL of a Telegram file, which is then used by the `DownloadFile` function to download the file. The `GetEnvsForSST` function retrieves environment variables for SST functionality.

**Unclear/Dead Code**
--------------------

No unclear or dead code was found in the provided files.

**