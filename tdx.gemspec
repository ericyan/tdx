Gem::Specification.new do |spec|
  spec.name        = 'tdx'
  spec.version     = '0.0.0'
  spec.author      = 'Eric Yan'
  spec.email       = 'long@ericyan.me'
  spec.homepage    = 'https://github.com/ericyan/tdx'
  spec.summary     = 'A library for parsing TDX market data feeds'
  spec.description = 'Parse market data feeds provided by TDX for historical prices of shares trading on Chinese stock exchanges.'
  spec.license     = 'MIT'

  spec.add_runtime_dependency 'bindata'
  spec.add_runtime_dependency 'time-series'

  spec.files       = `git ls-files -z`.split("\x0")

  spec.require_paths = ["lib"]
end
