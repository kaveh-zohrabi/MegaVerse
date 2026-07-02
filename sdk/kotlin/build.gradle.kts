plugins {
    kotlin("jvm") version "1.9.22"
}

group = "dev.megaverse"
version = "0.1.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation("com.squareup.okhttp3:okhttp:4.12.0")
    implementation("com.google.code.gson:gson:2.10.1")
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.7.3")
}

kotlin {
    jvmToolchain(21)
}
