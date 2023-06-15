// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package archive

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **NOTE**: This resource is deprecated, use data source instead.
type File struct {
	pulumi.CustomResourceState

	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
	ExcludeSymlinkDirectories pulumi.BoolPtrOutput `pulumi:"excludeSymlinkDirectories"`
	// Specify files to ignore when reading the `sourceDir`.
	Excludes pulumi.StringArrayOutput `pulumi:"excludes"`
	// Base64 Encoded SHA256 checksum of output file
	OutputBase64sha256 pulumi.StringOutput `pulumi:"outputBase64sha256"`
	// Base64 Encoded SHA512 checksum of output file
	OutputBase64sha512 pulumi.StringOutput `pulumi:"outputBase64sha512"`
	// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	OutputFileMode pulumi.StringPtrOutput `pulumi:"outputFileMode"`
	// MD5 of output file
	OutputMd5 pulumi.StringOutput `pulumi:"outputMd5"`
	// The output of the archive file.
	OutputPath pulumi.StringOutput `pulumi:"outputPath"`
	// SHA1 checksum of output file
	OutputSha pulumi.StringOutput `pulumi:"outputSha"`
	// SHA256 checksum of output file
	OutputSha256 pulumi.StringOutput `pulumi:"outputSha256"`
	// SHA512 checksum of output file
	OutputSha512 pulumi.StringOutput `pulumi:"outputSha512"`
	// The byte size of the output archive file.
	OutputSize pulumi.IntOutput `pulumi:"outputSize"`
	// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContent pulumi.StringPtrOutput `pulumi:"sourceContent"`
	// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContentFilename pulumi.StringPtrOutput `pulumi:"sourceContentFilename"`
	// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceDir pulumi.StringPtrOutput `pulumi:"sourceDir"`
	// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceFile pulumi.StringPtrOutput `pulumi:"sourceFile"`
	// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	Sources FileSourceArrayOutput `pulumi:"sources"`
	// The type of archive to generate. NOTE: `zip` is supported.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFile registers a new resource with the given unique name, arguments, and options.
func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OutputPath == nil {
		return nil, errors.New("invalid value for required argument 'OutputPath'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource File
	err := ctx.RegisterResource("archive:index/file:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFile gets an existing File resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("archive:index/file:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering File resources.
type fileState struct {
	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
	ExcludeSymlinkDirectories *bool `pulumi:"excludeSymlinkDirectories"`
	// Specify files to ignore when reading the `sourceDir`.
	Excludes []string `pulumi:"excludes"`
	// Base64 Encoded SHA256 checksum of output file
	OutputBase64sha256 *string `pulumi:"outputBase64sha256"`
	// Base64 Encoded SHA512 checksum of output file
	OutputBase64sha512 *string `pulumi:"outputBase64sha512"`
	// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	OutputFileMode *string `pulumi:"outputFileMode"`
	// MD5 of output file
	OutputMd5 *string `pulumi:"outputMd5"`
	// The output of the archive file.
	OutputPath *string `pulumi:"outputPath"`
	// SHA1 checksum of output file
	OutputSha *string `pulumi:"outputSha"`
	// SHA256 checksum of output file
	OutputSha256 *string `pulumi:"outputSha256"`
	// SHA512 checksum of output file
	OutputSha512 *string `pulumi:"outputSha512"`
	// The byte size of the output archive file.
	OutputSize *int `pulumi:"outputSize"`
	// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContent *string `pulumi:"sourceContent"`
	// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContentFilename *string `pulumi:"sourceContentFilename"`
	// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceDir *string `pulumi:"sourceDir"`
	// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceFile *string `pulumi:"sourceFile"`
	// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	Sources []FileSource `pulumi:"sources"`
	// The type of archive to generate. NOTE: `zip` is supported.
	Type *string `pulumi:"type"`
}

type FileState struct {
	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
	ExcludeSymlinkDirectories pulumi.BoolPtrInput
	// Specify files to ignore when reading the `sourceDir`.
	Excludes pulumi.StringArrayInput
	// Base64 Encoded SHA256 checksum of output file
	OutputBase64sha256 pulumi.StringPtrInput
	// Base64 Encoded SHA512 checksum of output file
	OutputBase64sha512 pulumi.StringPtrInput
	// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	OutputFileMode pulumi.StringPtrInput
	// MD5 of output file
	OutputMd5 pulumi.StringPtrInput
	// The output of the archive file.
	OutputPath pulumi.StringPtrInput
	// SHA1 checksum of output file
	OutputSha pulumi.StringPtrInput
	// SHA256 checksum of output file
	OutputSha256 pulumi.StringPtrInput
	// SHA512 checksum of output file
	OutputSha512 pulumi.StringPtrInput
	// The byte size of the output archive file.
	OutputSize pulumi.IntPtrInput
	// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContent pulumi.StringPtrInput
	// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContentFilename pulumi.StringPtrInput
	// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceDir pulumi.StringPtrInput
	// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceFile pulumi.StringPtrInput
	// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	Sources FileSourceArrayInput
	// The type of archive to generate. NOTE: `zip` is supported.
	Type pulumi.StringPtrInput
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
	ExcludeSymlinkDirectories *bool `pulumi:"excludeSymlinkDirectories"`
	// Specify files to ignore when reading the `sourceDir`.
	Excludes []string `pulumi:"excludes"`
	// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	OutputFileMode *string `pulumi:"outputFileMode"`
	// The output of the archive file.
	OutputPath string `pulumi:"outputPath"`
	// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContent *string `pulumi:"sourceContent"`
	// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContentFilename *string `pulumi:"sourceContentFilename"`
	// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceDir *string `pulumi:"sourceDir"`
	// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceFile *string `pulumi:"sourceFile"`
	// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	Sources []FileSource `pulumi:"sources"`
	// The type of archive to generate. NOTE: `zip` is supported.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a File resource.
type FileArgs struct {
	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
	ExcludeSymlinkDirectories pulumi.BoolPtrInput
	// Specify files to ignore when reading the `sourceDir`.
	Excludes pulumi.StringArrayInput
	// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	OutputFileMode pulumi.StringPtrInput
	// The output of the archive file.
	OutputPath pulumi.StringInput
	// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContent pulumi.StringPtrInput
	// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceContentFilename pulumi.StringPtrInput
	// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceDir pulumi.StringPtrInput
	// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	SourceFile pulumi.StringPtrInput
	// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
	Sources FileSourceArrayInput
	// The type of archive to generate. NOTE: `zip` is supported.
	Type pulumi.StringInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

// FileArrayInput is an input type that accepts FileArray and FileArrayOutput values.
// You can construct a concrete instance of `FileArrayInput` via:
//
//	FileArray{ FileArgs{...} }
type FileArrayInput interface {
	pulumi.Input

	ToFileArrayOutput() FileArrayOutput
	ToFileArrayOutputWithContext(context.Context) FileArrayOutput
}

type FileArray []FileInput

func (FileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (i FileArray) ToFileArrayOutput() FileArrayOutput {
	return i.ToFileArrayOutputWithContext(context.Background())
}

func (i FileArray) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileArrayOutput)
}

// FileMapInput is an input type that accepts FileMap and FileMapOutput values.
// You can construct a concrete instance of `FileMapInput` via:
//
//	FileMap{ "key": FileArgs{...} }
type FileMapInput interface {
	pulumi.Input

	ToFileMapOutput() FileMapOutput
	ToFileMapOutputWithContext(context.Context) FileMapOutput
}

type FileMap map[string]FileInput

func (FileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (i FileMap) ToFileMapOutput() FileMapOutput {
	return i.ToFileMapOutputWithContext(context.Background())
}

func (i FileMap) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileMapOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
func (o FileOutput) ExcludeSymlinkDirectories() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *File) pulumi.BoolPtrOutput { return v.ExcludeSymlinkDirectories }).(pulumi.BoolPtrOutput)
}

// Specify files to ignore when reading the `sourceDir`.
func (o FileOutput) Excludes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *File) pulumi.StringArrayOutput { return v.Excludes }).(pulumi.StringArrayOutput)
}

// Base64 Encoded SHA256 checksum of output file
func (o FileOutput) OutputBase64sha256() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputBase64sha256 }).(pulumi.StringOutput)
}

// Base64 Encoded SHA512 checksum of output file
func (o FileOutput) OutputBase64sha512() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputBase64sha512 }).(pulumi.StringOutput)
}

// String that specifies the octal file mode for all archived files. For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
func (o FileOutput) OutputFileMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.OutputFileMode }).(pulumi.StringPtrOutput)
}

// MD5 of output file
func (o FileOutput) OutputMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputMd5 }).(pulumi.StringOutput)
}

// The output of the archive file.
func (o FileOutput) OutputPath() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputPath }).(pulumi.StringOutput)
}

// SHA1 checksum of output file
func (o FileOutput) OutputSha() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputSha }).(pulumi.StringOutput)
}

// SHA256 checksum of output file
func (o FileOutput) OutputSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputSha256 }).(pulumi.StringOutput)
}

// SHA512 checksum of output file
func (o FileOutput) OutputSha512() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.OutputSha512 }).(pulumi.StringOutput)
}

// The byte size of the output archive file.
func (o FileOutput) OutputSize() pulumi.IntOutput {
	return o.ApplyT(func(v *File) pulumi.IntOutput { return v.OutputSize }).(pulumi.IntOutput)
}

// Add only this content to the archive with `sourceContentFilename` as the filename. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
func (o FileOutput) SourceContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.SourceContent }).(pulumi.StringPtrOutput)
}

// Set this as the filename when using `sourceContent`. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
func (o FileOutput) SourceContentFilename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.SourceContentFilename }).(pulumi.StringPtrOutput)
}

// Package entire contents of this directory into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
func (o FileOutput) SourceDir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.SourceDir }).(pulumi.StringPtrOutput)
}

// Package this file into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
func (o FileOutput) SourceFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.SourceFile }).(pulumi.StringPtrOutput)
}

// Specifies attributes of a single source file to include into the archive. One and only one of `source`, `sourceContentFilename` (with `sourceContent`), `sourceFile`, or `sourceDir` must be specified.
func (o FileOutput) Sources() FileSourceArrayOutput {
	return o.ApplyT(func(v *File) FileSourceArrayOutput { return v.Sources }).(FileSourceArrayOutput)
}

// The type of archive to generate. NOTE: `zip` is supported.
func (o FileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FileArrayOutput struct{ *pulumi.OutputState }

func (FileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*File)(nil)).Elem()
}

func (o FileArrayOutput) ToFileArrayOutput() FileArrayOutput {
	return o
}

func (o FileArrayOutput) ToFileArrayOutputWithContext(ctx context.Context) FileArrayOutput {
	return o
}

func (o FileArrayOutput) Index(i pulumi.IntInput) FileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *File {
		return vs[0].([]*File)[vs[1].(int)]
	}).(FileOutput)
}

type FileMapOutput struct{ *pulumi.OutputState }

func (FileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*File)(nil)).Elem()
}

func (o FileMapOutput) ToFileMapOutput() FileMapOutput {
	return o
}

func (o FileMapOutput) ToFileMapOutputWithContext(ctx context.Context) FileMapOutput {
	return o
}

func (o FileMapOutput) MapIndex(k pulumi.StringInput) FileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *File {
		return vs[0].(map[string]*File)[vs[1].(string)]
	}).(FileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileInput)(nil)).Elem(), &File{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileArrayInput)(nil)).Elem(), FileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileMapInput)(nil)).Elem(), FileMap{})
	pulumi.RegisterOutputType(FileOutput{})
	pulumi.RegisterOutputType(FileArrayOutput{})
	pulumi.RegisterOutputType(FileMapOutput{})
}
