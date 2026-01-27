# typed: false
# frozen_string_literal: true

class LinearCli < Formula
  desc "Token-efficient CLI for Linear"
  homepage "https://github.com/joa23/linear-cli"
  version "1.4.5"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_arm64.tar.gz"
      sha256 "bf86023119783b37f301968e2c60c7905353f33c045c046390e70bbd0f8afcd9"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Darwin_x86_64.tar.gz"
      sha256 "81e10b0a5bd5efefc69de6423b9d161b2ddb6db64e5883104ca2602d458cc102"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_arm64.tar.gz"
      sha256 "60d18f2708cae602b0dfd7e1cc9fc00dff1adbde93cf1d602bc76b4e54e4c979"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-cli_Linux_x86_64.tar.gz"
      sha256 "3c71f1081e88a53f188b236293ddd7ea390d98c0b0b98f097c3340aebc730021"
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
