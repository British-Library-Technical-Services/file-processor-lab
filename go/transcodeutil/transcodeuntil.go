package transcodeutil

import (
	"os/exec"
	"file-processor-lab/fileutil"
)

func FileTranscode(fd *fileutil.FileData) (string, error) {

	target := fd.File.Name() + ".mp3"

	cmd := exec.Command("ffmpeg", "-hide_banner", "-i", fd.File.Name(), "-c:a", "libmp3lame", "-b:a", "256k", "-ar", "44100", target)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string("Transcoding" + fd.File.Name()), nil
}