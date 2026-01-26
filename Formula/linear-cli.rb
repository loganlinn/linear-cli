# typed: false
# frozen_string_literal: true

class LinearCli < Formula
  desc "Token-efficient CLI for Linear"
  homepage "https://github.com/joa23/linear-cli"
  version "1.4.1"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_arm64.tar.gz"
      sha256 "582e1f6cdca8039ecf05c4bb7843d8c28e5cb3603cffd3c98138e9f3f489f544"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_x86_64.tar.gz"
      sha256 "420c78bb573fd408778269add93eedbb13f0f5a15c3dbf8e12e66b4dcb01c114"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_arm64.tar.gz"
      sha256 "86993c3409466262565c07196d2f08653df644ccc099ec7b1772329e2b83d2f3"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_x86_64.tar.gz"
      sha256 "35a76823ad1d30df1dbf20070e2cce54176c76a6f3d1d11c34bdd4981ce7d9fe"
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
