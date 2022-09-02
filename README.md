Test App made using Fyne

### Getting Started

1. Install fyne bundle tools

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

2. Bundle static assets

```bash
make bundle
```

3. Run locally

```bash
make run
```

### Compiling mobile app (android)

This can compile into an apk for android and ios devices if the corresponding SDKs are installed

Android:
Easiest to install via Android Studio SDK manager

- Android API 33
- Android 12 (S)
- NDK 21.2.6472646
  - Also set this in the path as the `ANDROID_NDK_HOME`
