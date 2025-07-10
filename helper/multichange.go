package helper

import (
	"bitmap/models"
	"bitmap/changetypes"
	"fmt"
	"strconv"
	"strings"
)

func ParseChanges(args []string) ([]models.Change, error) {
	var changes []models.Change

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--mirror="):
			direction := strings.TrimPrefix(arg, "--mirror=")
			direction = strings.ToLower(direction)

			mirrorOpts := models.MirrorOpts{
				Horizontal: strings.Contains("h,hor,horizontal,horizontally", direction),
				Vertical:   strings.Contains("v,ver,vertical,vertically", direction),
			}

			if !mirrorOpts.Horizontal && !mirrorOpts.Vertical {
				return nil, fmt.Errorf("invalid mirror direction: %s", direction)
			}

			changes = append(changes, models.Change{
				Type: models.ChangeTypeMirror,
				Args: []string{direction},
			})

		case strings.HasPrefix(arg, "--filter="):
			filter := strings.TrimPrefix(arg, "--filter=")
			filters := strings.Split(filter, ",")

			validFilters := map[string]bool{
				"red":       true,
				"green":     true,
				"blue":      true,
				"grayscale": true,
				"negative":  true,
				"pixelate":  true,
				"blur":      true,
			}

			for _, f := range filters {
				if !validFilters[f] {
					return nil, fmt.Errorf("invalid filter: %s", f)
				}
			}

			changes = append(changes, models.Change{
				Type: models.ChangeTypeFilter,
				Args: filters,
			})
		case strings.HasPrefix(arg, "--rotate="):

			rotation := strings.TrimPrefix(arg, "--rotate=")
			rotations := strings.Split(rotation, ",")
			changes = append(changes, models.Change{
				Type: models.ChangeTypeRotate,
				Args: rotations,
			})
		case strings.HasPrefix(arg, "--crop="):
			cropStr := strings.TrimPrefix(arg, "--crop=")
			cropValues := strings.Split(cropStr, "-")

			if len(cropValues) == 2 {
				cropValues = append(cropValues, "0", "0")
			}

			if len(cropValues) != 4 {
				return nil, fmt.Errorf("invalid crop format: expected 2 or 4 values, but got %d", len(cropValues))
			}

			offsetX, err := strconv.Atoi(cropValues[0])
			if err != nil {
				return nil, fmt.Errorf("invalid OffsetX value: %v", err)
			}
			offsetY, err := strconv.Atoi(cropValues[1])
			if err != nil {
				return nil, fmt.Errorf("invalid OffsetY value: %v", err)
			}
			width, err := strconv.Atoi(cropValues[2])
			if err != nil {
				return nil, fmt.Errorf("invalid Width value: %v", err)
			}
			height, err := strconv.Atoi(cropValues[3])
			if err != nil {
				return nil, fmt.Errorf("invalid Height value: %v", err)
			}

			changes = append(changes, models.Change{
				Type: models.ChangeTypeCrop,
				Args: []string{fmt.Sprintf("%d", offsetX), fmt.Sprintf("%d", offsetY), fmt.Sprintf("%d", width), fmt.Sprintf("%d", height)},
			})

		default:
			return nil, ErrInvalidArgs
		}
	}

	if len(changes) == 0 {
		return nil, fmt.Errorf("no valid change type found")
	}

	return changes, nil
}

func ApplyChanges(img *models.BMPImg, changes []models.Change) error {
	if len(changes) == 0 {
		return fmt.Errorf("no changes to apply")
	}

	for _, change := range changes {
		switch change.Type {
		case models.ChangeTypeMirror:
			if len(change.Args) == 0 {
				return ErrInvalidArgs
			}
			direction := change.Args[0]
			changetypes.Mirror(img, direction)
		case models.ChangeTypeFilter:
			if len(change.Args) == 0 {
				return ErrInvalidArgs
			}
			filters := change.Args
			changetypes.Filter(img, filters)

		case models.ChangeTypeRotate:

			for _, arg := range change.Args {
				arg = strings.ToLower(strings.TrimSpace(arg))

				var angle int
				switch arg {
				case "right", "90", "-270":
					angle = 90
				case "left", "-90", "270":
					angle = -90
				case "180", "-180":
					angle = 180
				case "360", "0":
					angle = 360
				default:
					return fmt.Errorf("invalid rotation direction: %s", arg)
				}

				changetypes.Rotate(img, angle)
			}
		case models.ChangeTypeCrop:
			if len(change.Args) != 4 {
				return fmt.Errorf("crop requires 4 arguments: OffsetX, OffsetY, Width, Height")
			}

			offsetX, err := strconv.Atoi(change.Args[0])
			if err != nil {
				return fmt.Errorf("invalid OffsetX: %v", err)
			}
			offsetY, err := strconv.Atoi(change.Args[1])
			if err != nil {
				return fmt.Errorf("invalid OffsetY: %v", err)
			}
			width, err := strconv.Atoi(change.Args[2])
			if err != nil {
				return fmt.Errorf("invalid Width: %v", err)
			}
			height, err := strconv.Atoi(change.Args[3])
			if err != nil {
				return fmt.Errorf("invalid Height: %v", err)
			}

			opts := models.CropOpts{
				OffsetX: offsetX,
				OffsetY: offsetY,
				Width:   width,
				Height:  height,
			}
			err = changetypes.Crop(img, opts)
			if err != nil {
				return fmt.Errorf("error applying crop: %v", err)
			}

		default:
			return fmt.Errorf("invalid change type: %d", change.Type)
		}
	}
	return nil
}
