# Panoptes

Panoptes is a mac screensaver that rotates through a set of webpages. Ideal for displaying monitoring dashboards as your screensaver.

## Configuration

Panoptes currently supports configuration via the command line.

### Configuring from the CLI

Panoptes can be configured from the command line.

#### List of URLs

Get the List of URLs like this:

```console
~ % defaults read ~/Library/Containers/com.apple.ScreenSaver.Engine.legacyScreenSaver/Data/Library/Preferences/ByHost/com.softwarepunk.Panoptes urls
(
    "https://www.apple.com"
)
~ %
```

Set the list of URLs:

```console
~ % defaults write ~/Library/Containers/com.apple.ScreenSaver.Engine.legacyScreenSaver/Data/Library/Preferences/ByHost/com.softwarepunk.Panoptes urls -array "https://www.apple.com" "https://www.google.com" 
~ %
```

#### Rotation Interval

The rotation interval is the time, in seconds, to display a URL before switching to a new URL.

Get it like this:

```console
~ % defaults read ~/Library/Containers/com.apple.ScreenSaver.Engine.legacyScreenSaver/Data/Library/Preferences/ByHost/com.softwarepunk.Panoptes intervalSecs
60
~ %
```

Set it like this:

```console
~ % defaults write ~/Library/Containers/com.apple.ScreenSaver.Engine.legacyScreenSaver/Data/Library/Preferences/ByHost/com.softwarepunk.Panoptes intervalSecs -float 30.0
~ %
```

