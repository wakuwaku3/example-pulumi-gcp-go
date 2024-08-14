package iac_storage

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BucketEnv interface {
	BucketSuffix() string
}

type Bucket struct {
	Name string
}

func (b *Bucket) UniqueName(env BucketEnv) string {
	return b.Name + "-" + env.BucketSuffix()
}

func (b *Bucket) ConvertPulumi(ctx *pulumi.Context, env BucketEnv, provider pulumi.ProviderResource) (*storage.Bucket, error) {
	return storage.NewBucket(
		ctx,
		b.Name,
		&storage.BucketArgs{
			Location: pulumi.String("ASIA-NORTHEAST1"),
			Name:     pulumi.String(b.UniqueName(env)),
		},
		pulumi.Provider(provider),
	)
}

var (
	Module1Bucket = &Bucket{Name: "my-bucket1"}
	Module2Bucket = &Bucket{Name: "my-bucket2"}
)
