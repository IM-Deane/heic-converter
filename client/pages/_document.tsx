import { Html, Head, Main, NextScript } from "next/document";
import siteConfig from "site.config";

export default function Document() {
	return (
		<Html lang="en" className="h-full bg-gray-50">
			<Head>
				<link
					rel="apple-touch-icon"
					sizes="180x180"
					href="/favicon/apple-touch-icon.png"
				/>
				<link
					rel="icon"
					type="image/png"
					sizes="32x32"
					href="/favicon/favicon-32x32.png"
				/>
				<link
					rel="icon"
					type="image/png"
					sizes="16x16"
					href="/favicon/favicon-16x16.png"
				/>
				<link
					rel="manifest"
					href="/manifest.json"
					crossOrigin="use-credentials"
				/>
				<link
					rel="mask-icon"
					href="/favicon/safari-pinned-tab.svg"
					color="#5bbad5"
				/>
				<meta name="msapplication-TileColor" content="#da532c" />
				<meta name="theme-color" content="#ffffff" />
				<meta name="description" content={siteConfig.description} />
			</Head>
			<body className="h-full overflow-hidden">
				<Main />
				<NextScript />
			</body>
		</Html>
	);
}
