# vector-go-sdk

A go SDK for Vector! This is an ***early alpha*** and not all features are present.  

## Usage

Currently. the SDK expects to find in your home directory the Python SDK data.
It consists of a .anki_vector folder where a sdk_config.ini file is placed. This file contains the configuration for each robot:

[ROBOT SERIAL NO]

cert = path to certificate (not needed by the go SDK)

ip = robot ip

name = robot name

guid = token for connecting to the robot

When you instantiate the robot you provide the serial number, then the SDK code will parse the .anki_vector/sdk_config.ini and get the data it needs.
I develop and test this version of the go SDK using a production Vector 1.0 and [wirepod](https://github.com/kercre123/wire-pod)

## Examples

Please see the cmd/examples directory for usage examples.

## Known Issues

- FaceEnrollmentStart(personName string)
  - Adding a new face enrollment via SDK doesn't seem to work. The procedure completes but the face is not saved on robot
- SaveHiResCameraPicture(fileName string) 
  - This doesn't seem to work on my production Vector on Wirepod. Probably vector-cloud on the robot needs to be updated. But since it's a production robot I'm stuck to 1.8...
    Anyway the image is saved at the regular 360p size.
- WriteText() / DisplayImage() / DisplayImageWithTransition()
  - Image and text displayed have garbled colors 

## Changelog 

***RELEASE_ALPHA_02***

- To try to support hires camera image snapshot I added OpenCV as a dependency. To install OpenCV on a Raspberry Pi or other Linux environment, run setup.sh
- Added camera functions
- Added face enrollment functions
- Added events with examples (Vector roars when touched)

***RELEASE_ALPHA_01***

I have taken the go SDK and did some changes.
- Integrated the functions of the great Wire's [vector-web-api-app](https://github.com/kercre123/vector-web-api-app) as sdk_wrapper functions
- Added many examples:
- The SDK supports features in these fields: 
  - Audio playback and volume settings
  - Custom eye colors
  - Animations
  - Settings
