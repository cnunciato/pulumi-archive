// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.archive.inputs;

import com.pulumi.archive.inputs.FileSourceArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FileState extends com.pulumi.resources.ResourceArgs {

    public static final FileState Empty = new FileState();

    /**
     * Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
     * 
     */
    @Import(name="excludeSymlinkDirectories")
    private @Nullable Output<Boolean> excludeSymlinkDirectories;

    /**
     * @return Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> excludeSymlinkDirectories() {
        return Optional.ofNullable(this.excludeSymlinkDirectories);
    }

    /**
     * Specify files to ignore when reading the `source_dir`.
     * 
     */
    @Import(name="excludes")
    private @Nullable Output<List<String>> excludes;

    /**
     * @return Specify files to ignore when reading the `source_dir`.
     * 
     */
    public Optional<Output<List<String>>> excludes() {
        return Optional.ofNullable(this.excludes);
    }

    /**
     * Base64 Encoded SHA256 checksum of output file
     * 
     */
    @Import(name="outputBase64sha256")
    private @Nullable Output<String> outputBase64sha256;

    /**
     * @return Base64 Encoded SHA256 checksum of output file
     * 
     */
    public Optional<Output<String>> outputBase64sha256() {
        return Optional.ofNullable(this.outputBase64sha256);
    }

    /**
     * Base64 Encoded SHA512 checksum of output file
     * 
     */
    @Import(name="outputBase64sha512")
    private @Nullable Output<String> outputBase64sha512;

    /**
     * @return Base64 Encoded SHA512 checksum of output file
     * 
     */
    public Optional<Output<String>> outputBase64sha512() {
        return Optional.ofNullable(this.outputBase64sha512);
    }

    /**
     * String that specifies the octal file mode for all archived files. For example: `&#34;0666&#34;`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
     * 
     */
    @Import(name="outputFileMode")
    private @Nullable Output<String> outputFileMode;

    /**
     * @return String that specifies the octal file mode for all archived files. For example: `&#34;0666&#34;`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
     * 
     */
    public Optional<Output<String>> outputFileMode() {
        return Optional.ofNullable(this.outputFileMode);
    }

    /**
     * MD5 of output file
     * 
     */
    @Import(name="outputMd5")
    private @Nullable Output<String> outputMd5;

    /**
     * @return MD5 of output file
     * 
     */
    public Optional<Output<String>> outputMd5() {
        return Optional.ofNullable(this.outputMd5);
    }

    /**
     * The output of the archive file.
     * 
     */
    @Import(name="outputPath")
    private @Nullable Output<String> outputPath;

    /**
     * @return The output of the archive file.
     * 
     */
    public Optional<Output<String>> outputPath() {
        return Optional.ofNullable(this.outputPath);
    }

    /**
     * SHA1 checksum of output file
     * 
     */
    @Import(name="outputSha")
    private @Nullable Output<String> outputSha;

    /**
     * @return SHA1 checksum of output file
     * 
     */
    public Optional<Output<String>> outputSha() {
        return Optional.ofNullable(this.outputSha);
    }

    /**
     * SHA256 checksum of output file
     * 
     */
    @Import(name="outputSha256")
    private @Nullable Output<String> outputSha256;

    /**
     * @return SHA256 checksum of output file
     * 
     */
    public Optional<Output<String>> outputSha256() {
        return Optional.ofNullable(this.outputSha256);
    }

    /**
     * SHA512 checksum of output file
     * 
     */
    @Import(name="outputSha512")
    private @Nullable Output<String> outputSha512;

    /**
     * @return SHA512 checksum of output file
     * 
     */
    public Optional<Output<String>> outputSha512() {
        return Optional.ofNullable(this.outputSha512);
    }

    /**
     * The byte size of the output archive file.
     * 
     */
    @Import(name="outputSize")
    private @Nullable Output<Integer> outputSize;

    /**
     * @return The byte size of the output archive file.
     * 
     */
    public Optional<Output<Integer>> outputSize() {
        return Optional.ofNullable(this.outputSize);
    }

    /**
     * Add only this content to the archive with `source_content_filename` as the filename. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    @Import(name="sourceContent")
    private @Nullable Output<String> sourceContent;

    /**
     * @return Add only this content to the archive with `source_content_filename` as the filename. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    public Optional<Output<String>> sourceContent() {
        return Optional.ofNullable(this.sourceContent);
    }

    /**
     * Set this as the filename when using `source_content`. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    @Import(name="sourceContentFilename")
    private @Nullable Output<String> sourceContentFilename;

    /**
     * @return Set this as the filename when using `source_content`. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    public Optional<Output<String>> sourceContentFilename() {
        return Optional.ofNullable(this.sourceContentFilename);
    }

    /**
     * Package entire contents of this directory into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    @Import(name="sourceDir")
    private @Nullable Output<String> sourceDir;

    /**
     * @return Package entire contents of this directory into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    public Optional<Output<String>> sourceDir() {
        return Optional.ofNullable(this.sourceDir);
    }

    /**
     * Package this file into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    @Import(name="sourceFile")
    private @Nullable Output<String> sourceFile;

    /**
     * @return Package this file into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    public Optional<Output<String>> sourceFile() {
        return Optional.ofNullable(this.sourceFile);
    }

    /**
     * Specifies attributes of a single source file to include into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    @Import(name="sources")
    private @Nullable Output<List<FileSourceArgs>> sources;

    /**
     * @return Specifies attributes of a single source file to include into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
     * 
     */
    public Optional<Output<List<FileSourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * The type of archive to generate. NOTE: `zip` is supported.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of archive to generate. NOTE: `zip` is supported.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private FileState() {}

    private FileState(FileState $) {
        this.excludeSymlinkDirectories = $.excludeSymlinkDirectories;
        this.excludes = $.excludes;
        this.outputBase64sha256 = $.outputBase64sha256;
        this.outputBase64sha512 = $.outputBase64sha512;
        this.outputFileMode = $.outputFileMode;
        this.outputMd5 = $.outputMd5;
        this.outputPath = $.outputPath;
        this.outputSha = $.outputSha;
        this.outputSha256 = $.outputSha256;
        this.outputSha512 = $.outputSha512;
        this.outputSize = $.outputSize;
        this.sourceContent = $.sourceContent;
        this.sourceContentFilename = $.sourceContentFilename;
        this.sourceDir = $.sourceDir;
        this.sourceFile = $.sourceFile;
        this.sources = $.sources;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FileState $;

        public Builder() {
            $ = new FileState();
        }

        public Builder(FileState defaults) {
            $ = new FileState(Objects.requireNonNull(defaults));
        }

        /**
         * @param excludeSymlinkDirectories Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder excludeSymlinkDirectories(@Nullable Output<Boolean> excludeSymlinkDirectories) {
            $.excludeSymlinkDirectories = excludeSymlinkDirectories;
            return this;
        }

        /**
         * @param excludeSymlinkDirectories Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder excludeSymlinkDirectories(Boolean excludeSymlinkDirectories) {
            return excludeSymlinkDirectories(Output.of(excludeSymlinkDirectories));
        }

        /**
         * @param excludes Specify files to ignore when reading the `source_dir`.
         * 
         * @return builder
         * 
         */
        public Builder excludes(@Nullable Output<List<String>> excludes) {
            $.excludes = excludes;
            return this;
        }

        /**
         * @param excludes Specify files to ignore when reading the `source_dir`.
         * 
         * @return builder
         * 
         */
        public Builder excludes(List<String> excludes) {
            return excludes(Output.of(excludes));
        }

        /**
         * @param excludes Specify files to ignore when reading the `source_dir`.
         * 
         * @return builder
         * 
         */
        public Builder excludes(String... excludes) {
            return excludes(List.of(excludes));
        }

        /**
         * @param outputBase64sha256 Base64 Encoded SHA256 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputBase64sha256(@Nullable Output<String> outputBase64sha256) {
            $.outputBase64sha256 = outputBase64sha256;
            return this;
        }

        /**
         * @param outputBase64sha256 Base64 Encoded SHA256 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputBase64sha256(String outputBase64sha256) {
            return outputBase64sha256(Output.of(outputBase64sha256));
        }

        /**
         * @param outputBase64sha512 Base64 Encoded SHA512 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputBase64sha512(@Nullable Output<String> outputBase64sha512) {
            $.outputBase64sha512 = outputBase64sha512;
            return this;
        }

        /**
         * @param outputBase64sha512 Base64 Encoded SHA512 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputBase64sha512(String outputBase64sha512) {
            return outputBase64sha512(Output.of(outputBase64sha512));
        }

        /**
         * @param outputFileMode String that specifies the octal file mode for all archived files. For example: `&#34;0666&#34;`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
         * 
         * @return builder
         * 
         */
        public Builder outputFileMode(@Nullable Output<String> outputFileMode) {
            $.outputFileMode = outputFileMode;
            return this;
        }

        /**
         * @param outputFileMode String that specifies the octal file mode for all archived files. For example: `&#34;0666&#34;`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
         * 
         * @return builder
         * 
         */
        public Builder outputFileMode(String outputFileMode) {
            return outputFileMode(Output.of(outputFileMode));
        }

        /**
         * @param outputMd5 MD5 of output file
         * 
         * @return builder
         * 
         */
        public Builder outputMd5(@Nullable Output<String> outputMd5) {
            $.outputMd5 = outputMd5;
            return this;
        }

        /**
         * @param outputMd5 MD5 of output file
         * 
         * @return builder
         * 
         */
        public Builder outputMd5(String outputMd5) {
            return outputMd5(Output.of(outputMd5));
        }

        /**
         * @param outputPath The output of the archive file.
         * 
         * @return builder
         * 
         */
        public Builder outputPath(@Nullable Output<String> outputPath) {
            $.outputPath = outputPath;
            return this;
        }

        /**
         * @param outputPath The output of the archive file.
         * 
         * @return builder
         * 
         */
        public Builder outputPath(String outputPath) {
            return outputPath(Output.of(outputPath));
        }

        /**
         * @param outputSha SHA1 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha(@Nullable Output<String> outputSha) {
            $.outputSha = outputSha;
            return this;
        }

        /**
         * @param outputSha SHA1 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha(String outputSha) {
            return outputSha(Output.of(outputSha));
        }

        /**
         * @param outputSha256 SHA256 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha256(@Nullable Output<String> outputSha256) {
            $.outputSha256 = outputSha256;
            return this;
        }

        /**
         * @param outputSha256 SHA256 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha256(String outputSha256) {
            return outputSha256(Output.of(outputSha256));
        }

        /**
         * @param outputSha512 SHA512 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha512(@Nullable Output<String> outputSha512) {
            $.outputSha512 = outputSha512;
            return this;
        }

        /**
         * @param outputSha512 SHA512 checksum of output file
         * 
         * @return builder
         * 
         */
        public Builder outputSha512(String outputSha512) {
            return outputSha512(Output.of(outputSha512));
        }

        /**
         * @param outputSize The byte size of the output archive file.
         * 
         * @return builder
         * 
         */
        public Builder outputSize(@Nullable Output<Integer> outputSize) {
            $.outputSize = outputSize;
            return this;
        }

        /**
         * @param outputSize The byte size of the output archive file.
         * 
         * @return builder
         * 
         */
        public Builder outputSize(Integer outputSize) {
            return outputSize(Output.of(outputSize));
        }

        /**
         * @param sourceContent Add only this content to the archive with `source_content_filename` as the filename. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceContent(@Nullable Output<String> sourceContent) {
            $.sourceContent = sourceContent;
            return this;
        }

        /**
         * @param sourceContent Add only this content to the archive with `source_content_filename` as the filename. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceContent(String sourceContent) {
            return sourceContent(Output.of(sourceContent));
        }

        /**
         * @param sourceContentFilename Set this as the filename when using `source_content`. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceContentFilename(@Nullable Output<String> sourceContentFilename) {
            $.sourceContentFilename = sourceContentFilename;
            return this;
        }

        /**
         * @param sourceContentFilename Set this as the filename when using `source_content`. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceContentFilename(String sourceContentFilename) {
            return sourceContentFilename(Output.of(sourceContentFilename));
        }

        /**
         * @param sourceDir Package entire contents of this directory into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceDir(@Nullable Output<String> sourceDir) {
            $.sourceDir = sourceDir;
            return this;
        }

        /**
         * @param sourceDir Package entire contents of this directory into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceDir(String sourceDir) {
            return sourceDir(Output.of(sourceDir));
        }

        /**
         * @param sourceFile Package this file into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceFile(@Nullable Output<String> sourceFile) {
            $.sourceFile = sourceFile;
            return this;
        }

        /**
         * @param sourceFile Package this file into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sourceFile(String sourceFile) {
            return sourceFile(Output.of(sourceFile));
        }

        /**
         * @param sources Specifies attributes of a single source file to include into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sources(@Nullable Output<List<FileSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Specifies attributes of a single source file to include into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<FileSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Specifies attributes of a single source file to include into the archive. One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder sources(FileSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param type The type of archive to generate. NOTE: `zip` is supported.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of archive to generate. NOTE: `zip` is supported.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public FileState build() {
            return $;
        }
    }

}
