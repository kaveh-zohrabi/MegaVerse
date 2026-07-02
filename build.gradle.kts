plugins {
    java
    kotlin("jvm") version "1.9.22"
    kotlin("plugin.spring") version "1.9.22"
}

group = "dev.megaverse"
version = "0.1.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
}

kotlin {
    jvmToolchain(21)
}
