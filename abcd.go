<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="六四记忆">
    <title>关于我们 - 六四记忆</title>
    <!-- Bootstrap 5 CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    <!-- Google Fonts -->
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@400;500;900&display=swap" rel="stylesheet">
    <!-- 自定义样式 -->
    <style>
        body {
            font-family: 'Noto Sans TC', sans-serif;
            background-color: #000000; /* 黑色背景 */
            color: #ffffff; /* 默认文字颜色为白色 */
            margin: 0;
            padding: 0;
            height: 100vh;
            display: flex;
            flex-direction: column; /* 列布局支持底部链接 */
            justify-content: space-between; /* 内容与底部分开 */
        }
        .intro-header {
            background-size: cover;
            background-position: center;
            width: 100%;
            padding: 40px 0; /* 上下内边距 */
        }
        .container {
            max-width: 800px; /* 限制内容宽度 */
        }
        .site-heading h1 {
            font-size: 2.5em;
            font-weight: 900;
            margin-bottom: 20px;
            color: #ffffff; /* 标题白色 */
        }
        .site-heading p {
            font-size: 1.1em;
            line-height: 1.5;
            margin-bottom: 15px;
        }
        .site-heading .subheading {
            display: block;
            font-size: 1em;
            margin-top: 20px;
            color: #ffffff; /* 子标题白色 */
        }
        .footer-link {
            position: fixed; /* 固定在页面底部 */
            bottom: 10px; /* 距离底部 10px */
            left: 10px; /* 距离左侧 10px，确保在左下角 */
            z-index: 1000; /* 确保在内容之上 */
        }
        .footer-link a {
            color: #ffffff; /* 白色文字 */
            font-weight: bold; /* 粗体 */
            text-decoration: underline; /* 下划线 */
            font-size: 1em;
        }
        .footer-link a:hover {
            color: #facc29; /* 悬停时变为黄色，与 index.html 高亮色一致 */
        }
    </style>
</head>
<body>
    <!-- Page Header -->
    <header class="intro-header">
        <div class="container">
            <div class="row">
                <div class="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
                    <div class="site-heading" id="tag-heading">
                        <h1>六四记忆</h1>
                        <p style="color: red; font-weight: bold;">
                            1989年6月4日，军队进入北京城，武力镇压持续近两个月的天安门民主运动，运动戛然而止。<br>
                            此后，官方不断抹除六四记忆。然而，30年来，作为这段历史的见证者，香港人以点燃30年<br>
                            的烛光和对中国民主运动的持续支持，践行承诺，守护记忆，捍卫人权，同时为香港的民主<br>
                            自由不懈奋斗。<br>
                            2020年，随着《港版国安法》迅速通过，悼念六四的烛光集会首次被定性为非法集会。<br>
                            当晚，香港支联会（香港市民支持爱国民主运动联合会）呼吁在法律危机下寻找永久空间保存<br>
                            历史记忆。<br>
                        </p>
                        <p style="color: red; font-weight: bold;">
                            本站以个人微薄之力，在网络上建立“六四记忆”，旨在<br>
                            唤起大陆良知，传播历史真相。<br>
                            我们希望号召更多人加入，共同守护真相，让历史不再被遗忘，让烛光继续照亮前行之路。
                        </p>
                        <span class="subheading">六四记忆</span>
                    </div>
                </div>
            </div>
        </div>
    </header>

    <!-- 返回主页链接 -->
    <div class="footer-link">
        <a href="/index.html">返回主页</a>
    </div>
</body>
</html>
