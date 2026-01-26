# typed: false
# frozen_string_literal: true

class LinearCli < Formula
  desc "Token-efficient CLI for Linear"
  homepage "https://github.com/joa23/linear-cli"
  version "1.3.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_arm64.tar.gz"
      sha256 "a7bda2426b39508120fed192f8c32767e52f9c6a107561c1de2ccd313fbefe02"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_x86_64.tar.gz"
      sha256 "317c7d6f61533a291b1a8b40bbc01065738f63bb5fab43f54cc64fe816188826"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_arm64.tar.gz"
      sha256 "ea1c261f92011f103e0631ada66930658f55542d1b0d7b2b21c528581bc18306"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_x86_64.tar.gz"
      sha256 "7b369021c1de5eb5600c01c87e41781ca9654190851c8809b33773b3808a2394"
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
