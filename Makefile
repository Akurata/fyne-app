
apk_name="DVD.apk"

.PHONY: bundle
bundle:
	@echo Bundling images...
	@fyne bundle \
		-o assets/bundled-images.go \
		--pkg assets \
		--prefix Asset \
		assets/img

	@echo Bundling fonts...
	@fyne bundle \
		-o assets/bundled-fonts.go \
		--pkg assets \
		--prefix Asset \
		assets/fonts

.PHONY: run
run:
	@go run .

.PHONY: android
android:
	@echo 1. Building Android bundle
	@fyne package -os android

	@echo 2. Installing app on device
	@adb install $(apk_name)