pluginManagement {
    repositories {
        mavenCentral()
        gradlePluginPortal()
    }
}

rootProject.name = "megaverse"

include(":services:social-service")
include(":services:crm-service")
include(":services:erp-service")
