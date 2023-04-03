import { siteConfig } from "@/types/site-config";
import { PhotoIcon } from "@heroicons/react/24/outline";

import logo from "./public/swift-convert.svg";

const domainName = process.env.NEXT_PUBLIC_DOMAIN_NAME;

export default siteConfig({
	siteName: "SwiftConvert",
	productBrand: logo,
	domain: domainName,
	developer: "Alchemized Software Ltd.",
	contactEmail: "hello@alchemizedsoftware.com",

	slogan: "SwiftConvert: Free, fast, and flawless HEIC conversions",
	description:
		"SwiftConvert is the free and fast solution for converting HEIC photos to JPEGs without sacrificing quality. Our tool ensures flawless and lightning-speed conversions, so you can easily share your photos without any hassle. Try SwiftConvert today and experience the convenience of hassle-free photo conversions.",

	mainNavTabs: [{ name: "Photos", href: "/", icon: PhotoIcon, current: true }],
});
