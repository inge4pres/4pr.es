package shortener

var PostHtml = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <title>4pres shortener</title>
        <meta name="description" content="growing knowledge one post at a time">
        <meta name="HandheldFriendly" content="True">
        <meta name="MobileOptimized" content="320">
        <meta name="generator" content="Hugo 0.17" />
        <meta name="robots" content="index,follow">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta property="og:title" content="4pres URL shortener">
        <meta property="og:description" content="growing knowledge one post at a time">
        <meta property="og:type" content="website">
        <meta property="og:url" content="https://4pr.es/">
        <link rel="stylesheet" href="https://inge.4pr.es/css/normalize.css">
        <link rel="stylesheet" href="https://inge.4pr.es/css/highlight.css">
        <link rel="stylesheet" href="https://inge.4pr.es/css/style.css">
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,400,600,700,300&subset=latin,cyrillic-ext,latin-ext,cyrillic">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">
    </head>
    <body>
     
<div id="wrapper">
        <header class="site-header">
            <div class="container">
                <div class="site-title-wrapper">
                    <h1 class="site-title">
                       <a title="4pres" href="/">4pres</a>
                    </h1>
					<a class="button-square button-social hint--top" data-hint="Github" title="Github" href="https://github.com/inge4pres/4pr.es/tree/serverless">
                        <i class="fa fa-github-alt"></i>
                    </a>
				</div>
			</div>
		</header>
		 <div id="container">


		<div id="post-index" class="container" itemscope="" itemtype="http://schema.org/Blog">
    
        	<header class="post-header">
        	    <h1 class="post-title" itemprop="name">4pres shortener</h1>
	       	     <p>there is your short url</p>
				<p>you can put it in social media posts, or keep it in your brain</p>
				<h2 class="post-stub" itemprop="surl">{{ .Url }}</h2>
        	</header>
		</div>
	</body>
</html>`

var Html404 = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <title>4pres shortener</title>
        <meta name="description" content="growing knowledge one post at a time">
        <meta name="HandheldFriendly" content="True">
        <meta name="MobileOptimized" content="320">
        <meta name="generator" content="Hugo 0.17" />
        <meta name="robots" content="index,follow">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta property="og:title" content="4pres URL shortener">
        <meta property="og:description" content="growing knowledge one post at a time">
        <meta property="og:type" content="website">
        <meta property="og:url" content="https://4pr.es/">
        <link rel="stylesheet" href="https://inge.4pr.es/css/normalize.css">
        <link rel="stylesheet" href="https://inge.4pr.es/css/highlight.css">
        <link rel="stylesheet" href="https://inge.4pr.es/css/style.css">
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Open+Sans:300italic,400italic,600italic,700italic,400,600,700,300&subset=latin,cyrillic-ext,latin-ext,cyrillic">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">
    </head>
    <body>
     
<div id="wrapper">
        <header class="site-header">
            <div class="container">
                <div class="site-title-wrapper">
                    <h1 class="site-title">
                       <a title="4pres" href="http://4pr.es/">4pres</a>
                    </h1>
					<a class="button-square button-social hint--top" data-hint="Github" title="Github" href="https://github.com/inge4pres/4pr.es/tree/serverless">
                        <i class="fa fa-github-alt"></i>
                    </a>
				</div>
			</div>
		</header>
		 <div id="container">


		<div id="post-index" class="container" itemscope="" itemtype="http://schema.org/Blog">
    
        <header class="post-header">
            <h1 class="post-title" itemprop="name">404 :(</h1>
            <p>you hit a blackhole bro! sorry... get a short URL <a href="http://4pr.es/">here</a></p>
			
        </header>
		
		</div>
	</body>
</html>`
