package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/wakuwaku3/example-pulumi-gcp-go/shared-module/env"
	"github.com/wakuwaku3/example-pulumi-gcp-go/shared-module/iac_storage"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		env, err := env.NewEnv()
		if err != nil {
			return err
		}

		provider, err := gcp.NewProvider(
			ctx,
			"gcp",
			&gcp.ProviderArgs{
				Project: pulumi.String(env.Project()),
				Region:  pulumi.String("asia-northeast1"),
			},
		)
		if err != nil {
			return err
		}

		// Create a Google Cloud resource (Storage Bucket)
		bucket, err := iac_storage.Module1Bucket.ConvertPulumi(
			ctx,
			env,
			provider,
		)
		if err != nil {
			return err
		}

		// Export the DNS name of the bucket
		ctx.Export("bucketName", bucket.Url)
		return nil
	})
}
