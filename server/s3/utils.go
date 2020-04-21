package s3

import "path"

func addResource(resources []*S3Resource, resource *S3Resource) []*S3Resource {
	if len(resource.parentDirs) == 0 {
		// will be adding/replacing in this resource array at this currentPath.
		for index, eResource := range resources {
			if eResource.Name == resource.Name {
				resources[index] = resource
				return resources
			}
		}
		resources = append(resources, resource)
		return resources
	}

	dirToFind := resource.parentDirs[0]

	// pass resource on to Dir resource.
	for index, eResource := range resources {
		if eResource.Name == dirToFind && eResource.Type == "Directory" {
			resource.traversePath()
			resources[index].Add(resource)
			return resources
		}
	}

	// otherwise create directory resource.
	resources = append(resources, &S3Resource{
		Name:       dirToFind,
		Type:       "Directory",
		BucketName: resource.BucketName,
		Path:       path.Join(resource.currentPath, dirToFind),
	})

	resources = addResource(resources, resource)
	return resources
}
