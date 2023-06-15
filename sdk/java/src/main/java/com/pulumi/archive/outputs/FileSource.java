// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.archive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FileSource {
    /**
     * @return Add this content to the archive with `filename` as the filename.
     * 
     */
    private String content;
    /**
     * @return Set this as the filename when declaring a `source`.
     * 
     */
    private String filename;

    private FileSource() {}
    /**
     * @return Add this content to the archive with `filename` as the filename.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return Set this as the filename when declaring a `source`.
     * 
     */
    public String filename() {
        return this.filename;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FileSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String content;
        private String filename;
        public Builder() {}
        public Builder(FileSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.filename = defaults.filename;
        }

        @CustomType.Setter
        public Builder content(String content) {
            this.content = Objects.requireNonNull(content);
            return this;
        }
        @CustomType.Setter
        public Builder filename(String filename) {
            this.filename = Objects.requireNonNull(filename);
            return this;
        }
        public FileSource build() {
            final var o = new FileSource();
            o.content = content;
            o.filename = filename;
            return o;
        }
    }
}
