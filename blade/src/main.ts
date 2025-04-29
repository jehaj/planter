import './style.css'
import typescriptLogo from '/typescript.svg'
import viteLogo from '/vite.svg'

const video = document.querySelector("video");
const constraints = {
  audio: false,
  video: true,
};

navigator.mediaDevices
  .getUserMedia(constraints)
  .then((stream) => {
    if (!video) return;
    const videoTracks = stream.getVideoTracks();
    const track = videoTracks[0];
    console.log("Got stream with constraints:", constraints);
    console.log(`Using video device: ${track.label}`);
    stream.onremovetrack = () => {
      console.log("Stream ended");
    };
    video.srcObject = stream;
  })
  .catch((error) => {
    if (!video) return;
    if (error.name === "OverconstrainedError") {
      console.error(
        `The resolution ${video.width}x${video.height} px is not supported by your device.`,
      );
    } else if (error.name === "NotAllowedError") {
      console.error(
        "You need to grant this page permission to access your camera and microphone.",
      );
    } else {
      console.error(`getUserMedia error: ${error.name}`, error);
    }
  });