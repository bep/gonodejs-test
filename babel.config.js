const presets = [
    [
      "@babel/env",
      {
        targets: {
            chrome: 52,
            browsers: ["last 20 versions", "safari 7"]
        },
      },
    ],
  ];
  
  module.exports = { presets };