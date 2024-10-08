# Octopusdeploy Resource Provider

**IMPORTANT**: Provider not ready. WIP!

The Octopus Deploy Resource Provider lets you manage [Octopus Deploy](http://example.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/octopusdeploy
```

or `yarn`:

```bash
yarn add @pulumi/octopusdeploy
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_octopus_deploy
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-octopus_deploy/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Octopusdeploy
```

## Configuration

The following configuration points are available for the `octopus_deploy` provider:

- `octopus_deploy:apiKey` (environment: `FOO_API_KEY`) - the API key for `octopus_deploy`
- `octopus_deploy:region` (environment: `FOO_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/octopus_deploy/api-docs/).
