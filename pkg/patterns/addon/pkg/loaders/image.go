package loaders

import "context"

// ImageRepository supports loading from Docker v2-2 images.
type ImageRepository struct {
	imageRef string
}

func (r *ImageRepository) LoadChannel(ctx context.Context, name string) (*Channel, error) {
	// TODO(njhale): Implement me
	return nil, nil
}

func (r *ImageRepository) LoadManifest(ctx context.Context, packageName, id string) (map[string]string, error) {
	// TODO(njhale): Implment me
	return nil, nil
}
