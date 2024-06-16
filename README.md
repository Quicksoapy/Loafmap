# Loafmap

Loafmap is a map of loaf stickers, where authenticated users can add a marker on the map, with a picture and a description of the loaf sticker they added anywhere in the world.

## Requirements

Docker Compose

## Set Secrets

```bash
cd Loafmap
```

```bash
cp --no-clobber --recursive --no-target-directory .secrets.example .secrets
```

## Usage

Run the application with a config using the `--config` or `-c` flag like so:

```bash
loafmap --config ./config.example.json
```
