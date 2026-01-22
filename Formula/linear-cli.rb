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
      sha256 "96353f93007eaf4aa837dacdfc5efee0d0c90e78ef09324b7f9d116e5d49a35e"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-darwin-amd64.tar.gz"
      sha256 "39b14b3c21db1575a30457b5f7ad101dd93b0604183dc7777a8011b5cea394ce"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-arm64.tar.gz"
      sha256 "53caf0067706b9938d0c9ea04d306349fae3a2507a317bccc81c6bbdfab4bdd7"
    end
    on_intel do
      url "https://github.com/joa23/linear-cli/releases/download/v#{version}/linear-linux-amd64.tar.gz"
      sha256 "8ce39aee5b7519c320096af3c46a62f1e34fbdd71f63270933c04090c77f8efd"
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
