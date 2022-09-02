
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

#
# Packaging helpers
#
.PHONY: package-android
package-android:
	@echo Building Android package
	@fyne package -os android

.PHONY: package-windows
package-windows:
	@echo Building Windows package
	@fyne package -os windows 

.PHONY: package-darwin
package-darwin:
	@echo Building MacOS package
	@fyne package -os darwin

.PHONY: package-ios
package-ios:
	@echo Building IOS package
	@fyne package -os ios 

#
# Compile and send to connected android device
#
.PHONY: android
android: package-android
	@echo Installing Android app on device
	@adb install $(apk_name)


#
# Prepare all files for release
#
.PHONY: release
release: bundle package-android package-windows