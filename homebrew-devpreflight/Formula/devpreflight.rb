class Devpreflight < Formula
  desc "CLI tool for validating development configurations and best practices"
  homepage "https://github.com/devpreflight/devpreflight"
  version "0.1.0" # This will be updated by GoReleaser

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/devpreflight/devpreflight/releases/download/v#{version}/devpreflight_Darwin_x86_64.tar.gz"
      sha256 "{{.SHA256}}" # This will be updated by GoReleaser
    else
      url "https://github.com/devpreflight/devpreflight/releases/download/v#{version}/devpreflight_Darwin_arm64.tar.gz"
      sha256 "{{.SHA256}}" # This will be updated by GoReleaser
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/devpreflight/devpreflight/releases/download/v#{version}/devpreflight_Linux_x86_64.tar.gz"
      sha256 "{{.SHA256}}" # This will be updated by GoReleaser
    else
      url "https://github.com/devpreflight/devpreflight/releases/download/v#{version}/devpreflight_Linux_arm64.tar.gz"
      sha256 "{{.SHA256}}" # This will be updated by GoReleaser
    end
  end

  def install
    bin.install "devpreflight"

    # Install shell completion
    output = Utils.safe_popen_read("#{bin}/devpreflight", "completion", "bash")
    (bash_completion/"devpreflight").write output

    output = Utils.safe_popen_read("#{bin}/devpreflight", "completion", "zsh")
    (zsh_completion/"_devpreflight").write output

    output = Utils.safe_popen_read("#{bin}/devpreflight", "completion", "fish")
    (fish_completion/"devpreflight.fish").write output
  end

  test do
    system "#{bin}/devpreflight", "version"
  end
end