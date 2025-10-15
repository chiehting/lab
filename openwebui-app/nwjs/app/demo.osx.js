import nwbuild from '../src/index.js';

await nwbuild({
  mode: 'build',
  flavor: 'sdk',
  platform: 'osx',
  srcDir: './app/src',
  cacheDir: './node_modules/nw',
  outDir: './app/out/osx',
  glob: false,
  logLevel: 'debug',
  app: {
    name: 'Demo',
    /* File path of icon from where it is copied. */
    icon: './app/src/icon.icns',
    LSApplicationCategoryType: 'public.app-category.utilities',
    CFBundleIdentifier: 'io.nwutils.demo',
    CFBundleName: 'Demo',
    CFBundleDisplayName: 'Demo',
    CFBundleSpokenName: 'Demo',
    CFBundleVersion: '0.0.0',
    CFBundleShortVersionString: '0.0.0',
    NSHumanReadableCopyright: 'Copyright (c) 2024 NW.js Utilities',
    NSLocalNetworkUsageDescription: 'Demo requires access to network to showcase its capabilities',
  }
});

console.log('\nExecute `npm run demo:exe:osx` to run the application.');
