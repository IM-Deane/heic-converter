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

	description: "A really good HEIF(.heic) and HEVC media transform.",

	// main navigation tabs
	mainNavTabs: [{ name: "Photos", href: "/", icon: PhotoIcon, current: true }],
});
