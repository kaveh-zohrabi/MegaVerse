# swift-tools-version: 5.9

import PackageDescription

let package = Package(
    name: "MegaVerseSDK",
    platforms: [
        .iOS(.v16),
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "MegaVerseSDK",
            targets: ["MegaVerseSDK"]
        ),
    ],
    dependencies: [
        .package(url: "https://github.com/Alamofire/Alamofire.git", from: "5.8.0"),
    ],
    targets: [
        .target(
            name: "MegaVerseSDK",
            dependencies: ["Alamofire"]
        ),
        .testTarget(
            name: "MegaVerseSDKTests",
            dependencies: ["MegaVerseSDK"]
        ),
    ]
)
