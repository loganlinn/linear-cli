# typed: false
# frozen_string_literal: true

class LinearCli < Formula
  desc "Token-efficient CLI for Linear"
  homepage "https://github.com/joa23/linear-cli"
  version "1.1.1"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-darwin-arm64.tar.gz"
      sha256 "d5c63ba5f18a9c060007dd1971abaa655e7a77d936b1bf367e423605d8383a51"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-darwin-amd64.tar.gz"
      sha256 "b2d28e367ab8f89d8b9fadd3a6f3e50441baf9faf70016542ca106c384088a68"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-arm64.tar.gz"
      sha256 "33c568e03cb32545edc546c92fb905ffbe162629a623afd07decae76e1ba48ff"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-amd64.tar.gz"
      sha256 "3d0cba41605ca748feb8e0e7339509bc4e01303f6bb294d72d743e42fb50c06e"
    end
  end

  # This is a binary distribution - no build step required
  def pour_bottle?
    true
  end

  def install
    bin.install "linear"
  end

  test do
    assert_match "Linear CLI", shell_output("#{bin}/linear --help")
  end
end
