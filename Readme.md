# This is the final project for ITIS 6177 Systems Integration at UNCC

endpoint:
https://eastus.api.cognitive.microsoft.com/

API Key set in .bashrc
## 4/17/24:

- made a lot of progress

- need to fix issue with text encoding from web form interface. spaces in text input need to be handled properly.

## 4/14/24:

- installed Tailwind and set up configuration, seems to be working!

- spent a lot of time troubleshooting with Air and then Templ. About 3 hours total. But after slogging through that I set up my templates nicely, switched over and very happy so far with everything.  It's like playing with React but faster and simpler and more intuitive.

## 4/13/24:
- Successfully took a GET request from a client containing text, obtained audio from TTS service and then returned an audio file to the client!

- Postman doesn't seem to be able to play files directly but the browser will play them. Next I'll need to allow the user to choose a voice from the list of voices as well as entering text.

- :8080/docs will now allow user to request voices list. next up, include ability to select voice and then enter text from docs page. 

