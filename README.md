# FiascoExtension

This program attempts to extend the encoding capabilities of the [FIASCO](https://github.com/l-tamas/Fiasco) image and
video codec by sharing references to domain images between multiple frames of a video file.

It was created in the course of the bachelor thesis "Erweiterung eines
fraktalen Videokompressionsverfahrens durch die Referenz derselben Fraktale"
(Extension of a fractal video compression method by the reference of the same
fractals), which can also be found in this repository.

## Usage

### Building the Project

```shell
go build
```

### Basic Encoding

```shell
FiascoExtension -a encode -i video.mp4 -o encoded.fco
```

### Basic Decoding

```shell
FiascoExtension -a decode -i encoded.fco -o video.mp4
```

## Requirements

The following prerequisites have to be met, to run this program successfully.

### Go

A Golang recent Golang installation.

### FFmpeg

A FFmpeg installation at version 4.3 or higher.

### FIASCO

The cfiaco and dfiasco binaries at version 1.3. They can be built from [https://github.com/l-tamas/Fiasco]().
You may have to tweak the source codes constants to avoid crashes.
For guidance on this see chapter 5.1.3 in the included thesis.

## Parameters

- `-a | --action`
    - One of: `[encode, decode]`. Decides if the program should encode a video file to the FIASCO codec of if it should
      decode an encoded FIASCO file.

- `-i | --input`
    - In case of encoding, the path to the video file that should be encoded. All file formats that are recognized by
      your FFmpeg version are applicable here.
    - In case of decoding, the path to the encoded FIASCO file.

- `-o | --output`
    - In case of encoding, the path where the encoded FIASCO file should be written to.
    - In case of decoding, the path where the decoded file should be written to. All file formats that are recognized by
      your FFmpeg version are applicable here.

- `-l | --layout`
    - Overrides the tiling layout used to tile videos to groups of pictures during encoding.
    - Overrides the tiling layout used to untile groups of pictures back into a video file during decoding.
    - Higher values result in longer encoding times, but better visual quality and compression.
    - Format: `\d+x\d+`, where the first number is the width, and the second number is the height of the tiling layout.
    - Default value: `1x8`.

- `-f | --fps`
    - Overrides the target fps of a decoded video file.
    - Default value: `25`.

- `ffmpegPath`
    - Overrides the path that the ffmpeg binary is called from during encoding and decoding.
    - Default value: `ffmpeg`.

- `cfiascoPath`
    - Overrides the path that the cfiasco binary is called from during encoding.
    - Default value: `cfiasco`.

- `dfiascoPath`
    - Overrides the path that the dfiasco binary is called from during decoding.
    - Default value: `dfiasco`.

- `--ffmpegArgs`
    - Additional arguments to be passed during the FFmpeg call when encoding of decoding.
    - Appended directly after the binary as-is.
    - See [FFmpeg docomentation](https://ffmpeg.org/ffmpeg.html) for details.

- `--fiascoArgs`
    - Additional arguments to be passed during the FIASCO call when encoding of decoding.
    - Appended directly after the binary as-is.
    - See FIASCO man-pages for details.
