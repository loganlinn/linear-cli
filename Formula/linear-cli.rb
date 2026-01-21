# typed: false
# frozen_string_literal: true

class LinearCli < Formula
  desc "Token-efficient CLI for Linear"
  homepage "https://github.com/joa23/linear-cli"
  version "1.0.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-darwin-arm64.tar.gz"
      sha256 "7713fb79c4164ed25d710f60daba062bd0f1d9252141937367e9caf0a19c562e"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-darwin-amd64.tar.gz"
      sha256 "021faa306e984ea407bbd9e2ca0a3c0aace5677d3f7ca531635a751b1db28cda"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-arm64.tar.gz"
      sha256 "a3a516f3c518db417bcc8b9d8aead90e5c8ae8ce1fef83199efe1da89dc67047"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-amd64.tar.gz"
      sha256 "f63869c117a5414176bd3df3fad4d85f5d0aac42aeb94778834a7d7c9f7e4557"
    end
  end

  def install
    bin.install "linear"
  end

  test do
    assert_match "Linear CLI", shell_output("#{bin}/linear --help")
  end
end
