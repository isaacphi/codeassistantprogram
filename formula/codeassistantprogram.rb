class Cap < Formula
  desc "Cap Code Assistant Program"
  homepage "https://github.com/isaacphi/codeassistantprogram"
  url "https://github.com/isaacphi/codeassistantprogram/archive/refs/tags/v0.1.0.tar.gz"
  sha256 "24d4204e92c47779859bd04487846a31afe56973bbaac675b32fee61ef9694bd"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "cap", *std_go_args(ldflags: "-s -w")
    bin.install "cap"
  end

  test do
    system "#{bin}/cap", "--version"
  end
end
