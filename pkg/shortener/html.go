package shortener

var PostHtml = `<!DOCTYPE html>
<html lang='en'>

<head>
  <meta charset='utf-8'>
  <meta name='viewport' content='width=device-width, initial-scale=1'>
  <meta name='description' content='rational thoughts'>

  <meta property='og:title' content='4pres'>
  <meta property='og:description' content=''>
  <meta property='og:url' content='https://4pr.es/'>
  <meta property='og:site_name' content='4pres'>
  <meta property='og:type' content='website'>

  <title>inge4pres</title>
  <link rel='canonical' href='https://inge.4pr.es/'>

  <link rel='icon' href='https://inge.4pr.es/favicon.ico'>
  <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu:400,400i,700&subset=latin'>
  <link rel='stylesheet' href='https://inge.4pr.es/assets/css/main.77da63e1.css'>
  <link rel='stylesheet' href='https://inge.4pr.es/css/custom.css'>
  <script>
    window.ga = window.ga || function () { (ga.q = ga.q || []).push(arguments) }; ga.l = +new Date;
    ga('create', 'UA-55252403-3', 'auto');
    ga('send', 'pageview');
  </script>
  <script async src='//www.google-analytics.com/analytics.js'></script>

</head>


<body class='home type-page'>
  <div class='site'>

    <header id='header' class='header-container'>
      <div class='header site-header'>
        <div class='header-info'>
          <h1 class='site-title title'>4PRES</h1>

          <p class='site-description subtitle'>Get a short URL for a long one</p>

        </div>
      </div>
    </header>

    <main id='main' class='main'>
      <div class='home-sections-container'>
        <div class='home-sections'>
          <section id='recent-posts' class='home-section'>
            <header>
                <p class='site-description'>Your short url is</p>
                <h2 class='site-title title' itemprop='surl'>{{ .Url }}</h2>          
                <p class='site-description subtitle'>you can put it in social media posts, or keep it for yourself</p>
            </header>
          </section>
        </div>
      </div>
    </main>

    <footer id='footer' class='footer-container'>
      <div class='footer'>
        <div class='social-menu-container'>
          <nav aria-label='Social Menu'>
            <ul class='social-menu'>
              <li>
                <a href='mailto:fgualazzi@gmail.com' target='_blank' rel='noopener'>
                  <span class='screen-reader'>Contact via Email</span>
                  <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                    <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z" />
                    <polyline points="22,6 12,13 2,6" />

                  </svg>
                </a>
              </li>
              <li>
                <a href='https://github.com/inge4pres/4pr.es/tree/serverless' target='_blank' rel='noopener'>
                  <span class='screen-reader'>Open Github account in new tab</span>
                  <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                    <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"
                    />

                  </svg>
                </a>
              </li>
              <li>
                <a href='https://twitter.com/inge4pres' target='_blank' rel='noopener'>
                  <span class='screen-reader'>Open Twitter account in new tab</span>
                  <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                    <path d="M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"
                    />

                  </svg>
                </a>
              </li>
            </ul>
          </nav>
        </div>
        <div class='copyright'>
          <p>
            inge4pres &copy; 2016
            <script>new Date().getFullYear() > 2016 && document.write("-" + new Date().getFullYear());</script> Francesco Gualazzi
          </p>
        </div>

      </div>
    </footer>

  </div>
  <script src='https://inge.4pr.es/assets/js/main.f29c93b9.js'></script>
  <script src='https://inge.4pr.es/js/custom.js'></script>
</body>

</html>`

var PostOld = `<!DOCTYPE html>
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
		<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-55252403-3', 'auto');
  ga('send', 'pageview');

</script>
</body>
</html>`

var Html404 = `<!DOCTYPE html>
<html lang='en'>

<head>
    <meta charset='utf-8'>
    <meta name='viewport' content='width=device-width, initial-scale=1'>
    <meta name='description' content='Growing knowledge one post at a time'>

    <meta property='og:title' content='Got lost? » inge4pres rational thoughts'>
    <meta property='og:description' content='Growing knowledge one post at a time'>
    <meta property='og:url' content=''>
    <meta property='og:site_name' content='inge4pres rational thoughts'>
    <meta property='og:type' content='website'>
    <meta property='og:updated_time' content='2017-10-14T15:27:36&#43;02:00' />
    <meta property='fb:app_id' content='1845367892442473'>
    <meta property='fb:admins' content='francesco.inge'>
    <meta name='twitter:card' content='summary'>

    <title>Got lost? » 4pres</title>
    <link rel='canonical' href=''>

    <link rel='icon' href='https://inge.4pr.es/favicon.ico'>
    <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu:400,400i,700&subset=latin'>
    <link rel='stylesheet' href='https://inge.4pr.es/assets/css/main.77da63e1.css'>
    <script>
        window.ga = window.ga || function () { (ga.q = ga.q || []).push(arguments) }; ga.l = +new Date;
        ga('create', 'UA-55252403-3', 'auto');
        ga('send', 'pageview');
    </script>
    <script async src='//www.google-analytics.com/analytics.js'></script>

</head>


<body class='_404 type-page'>
    <div class='site'>
        <header id='header' class='header-container'>
            <div class='header site-header'>
                <nav id='main-menu' class='main-menu-container' aria-label='Main Menu'>
                    <ul class='main-menu'>
                        <li>
                            <a href='/'>/</a>
                        </li>
                        <li>
                            <a href='/about'>/about</a>
                        </li>
                        <li>
                            <a href='/resume'>/resume</a>
                        </li>

                    </ul>
                </nav>
                <div class='header-info'>
                    <h1 class='site-title title'>4PRES</h1>
                    <p class='site-description subtitle'>Get a short URL for a long one</p>
                </div>
            </div>
        </header>

        <main id='main' class='main'>
            <div class='entry'>
                <div class='entry-content'>
                    <span class='screen-reader'>Got lost?</span>
                    <figure class='gopher'>
                        <a href='/'>
                            <img src='https://inge.4pr.es/gopher.png' alt='Gopher' />
                            <span class='screen-reader'>Go Home...</span>
                        </a>
                        <figcaption>
                            <h1 class='title'>
                                <a href='/'>Got lost?</a>
                            </h1>
                        </figcaption>
                    </figure>
                </div>
            </div>
        </main>

        <footer id='footer' class='footer-container'>
            <div class='footer'>
                <div class='social-menu-container'>
                    <nav aria-label='Social Menu'>
                        <ul class='social-menu'>
                            <li>
                                <a href='mailto:fgualazzi@gmail.com' target='_blank' rel='noopener'>
                                    <span class='screen-reader'>Contact via Email</span>
                                    <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                                        <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z" />
                                        <polyline points="22,6 12,13 2,6" />

                                    </svg>
                                </a>
                            </li>
                            <li>
                                <a href='https://github.com/inge4pres/4pr.es/tree/serverless' target='_blank' rel='noopener'>
                                    <span class='screen-reader'>Open Github account in new tab</span>
                                    <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                                        <path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"
                                        />

                                    </svg>
                                </a>
                            </li>
                            <li>
                                <a href='https://twitter.com/inge4pres' target='_blank' rel='noopener'>
                                    <span class='screen-reader'>Open Twitter account in new tab</span>
                                    <svg class='icon' viewbox='0 0 24 24' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' aria-hidden='true'>

                                        <path d="M23 3a10.9 10.9 0 0 1-3.14 1.53 4.48 4.48 0 0 0-7.86 3v1A10.66 10.66 0 0 1 3 4s-4 9 5 13a11.64 11.64 0 0 1-7 2c9 5 20 0 20-11.5a4.5 4.5 0 0 0-.08-.83A7.72 7.72 0 0 0 23 3z"
                                        />

                                    </svg>
                                </a>
                            </li>
                        </ul>
                    </nav>
                </div>
                <div class='copyright'>
                    <p>
                        inge4pres &copy; 2016
                        <script>new Date().getFullYear() > 2016 && document.write("-" + new Date().getFullYear());</script> Francesco Gualazzi
                    </p>
                </div>

            </div>
        </footer>

    </div>
    <script src='/assets/js/main.f29c93b9.js'></script>
</body>

</html>`
