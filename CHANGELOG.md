# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Implemented `ietf-tools/semver-action@v1` for robust semantic versioning
- Enhanced conventional commit support with comprehensive commit types
- Automatic version determination based on commit message analysis
- Improved release notes with version type and previous version information
- Separated GitHub Actions CI/CD workflows for each fx module
- Individual test and release workflows for fxConfig and fxEcho modules
- Module-level workflows for overall project changes
- Independent versioning for each module (fxconfig-v1.2.3, fxecho-v1.2.3, v1.2.3)
- Path-based workflow triggers for efficient CI/CD execution
- Comprehensive documentation for separated workflow structure

### Changed
- Updated README.md with detailed semver action documentation
- Enhanced commit convention documentation with full conventional commit support
- Improved project structure documentation to reflect new workflow organization
- Enhanced versioning strategy documentation

## [1.0.0] - 2025-01-XX

### Added
- Initial release of the UTOL Go module
- fxConfig module for configuration management using Uber's fx framework
- fxEcho module for Echo web framework integration with fx
- Viper-based configuration system with environment variable support
- Comprehensive test coverage for all modules
- Example configuration files and documentation
- GitHub Actions CI/CD workflows for automated testing and releases
- Comprehensive documentation for commit conventions and release process
- Automated semantic versioning based on commit message prefixes

### Features
- **fxConfig Module**:
  - Environment-based configuration loading
  - Support for YAML configuration files
  - Environment variable overrides
  - Hot-reload capability for configuration changes
  
- **fxEcho Module**:
  - Echo web framework integration
  - Dependency injection ready
  - Modular architecture for easy extension

[Unreleased]: https://github.com/UTOL-s/module/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/UTOL-s/module/releases/tag/v1.0.0 