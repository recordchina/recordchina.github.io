<!DOCTYPE html>
<html lang="zh-TW">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="记录历史 - 图片展示">
    <title>记录历史 - 永不忘</title>
    <!-- Bootstrap 5 CDN -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    <!-- Google Fonts -->
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@400;500;900&display=swap" rel="stylesheet">
    <!-- 自定义样式 -->
    <style>
        body {
            font-family: 'Noto Sans TC', sans-serif;
            background-color: #575f4f;
            color: #f1e8dd;
            margin: 0;
            height: 100vh;
            display: flex;
            flex-direction: column;
        }
        .header-mobile {
            display: flex;
            justify-content: space-between;
            padding: 15px;
        }
        .main {
            padding: 20px;
            flex-grow: 1;
            display: flex;
            align-items: center;
        }
        .left-column {
            height: 100%;
            display: flex;
            flex-direction: column;
            justify-content: center;
            padding-right: 10px; /* 减少右边距，靠近图片 */
        }
        .title h2 {
            font-size: 2em;
            font-weight: 900;
            margin: 0 0 10px 0;
        }
        .intro p {
            font-size: 1.1em;
            line-height: 1.65em;
            margin-bottom: 10px;
        }
        .cate {
            margin-top: 5px; /* 减少上边距，靠近图片 */
        }
        .cate ul {
            list-style: none;
            padding: 0;
            margin: 0;
            text-align: center; /* 年份居中对齐 */
        }
        .cate ul li a {
            color: #ff6b6b; /* 鲜艳珊瑚红 */
            text-decoration: none;
            padding: 5px 10px;
            display: block;
            font-size: 1.5em; /* 保持大字体 */
            font-weight: 500; /* 加粗 */
        }
        .cate ul li a:hover, .cate ul li a.active {
            color: #facc29; /* 悬停和选中保持原有活跃颜色 */
        }
        .content {
            height: 100%;
            display: flex;
            flex-direction: column;
            justify-content: center;
        }
        .content img {
            width: 100%;
            height: 100%;
            max-height: 507px;
            display: block;
            margin-bottom: 15px;
            border: 2px solid #f1e8dd;
            cursor: pointer;
        }
        .content .col-md-6 {
            padding: 5px;
            height: 507px;
        }
        .content .row {
            margin: 0;
        }
        .pagination {
            margin-top: 15px;
            display: flex;
            align-items: center;
            justify-content: center;
            flex-wrap: wrap;
        }
        .pagination button, .pagination input {
            background-color: #f1e8dd;
            color: #575f4f;
            border: none;
            padding: 5px 10px;
            margin: 0 5px;
            cursor: pointer;
        }
        .pagination button:disabled {
            background-color: #ccc;
            cursor: not-allowed;
        }
        .pagination button.active {
            background-color: #facc29;
            color: #575f4f;
        }
        .pagination input {
            width: 50px;
            text-align: center;
        }
        .footer {
            padding: 30px 15px;
            color: #fff;
        }
        /* 纪念文字容器 */
        .memory-text {
            text-align: left; /* 左对齐 */
            margin-bottom: 20px;
            padding: 15px;
            background: rgba(241, 232, 221, 0.2); /* 半透明背景 */
            border-radius: 5px;
        }
        .memory-text p {
            color: #facc29; /* 活跃颜色 */
            font-family: 'Noto Sans TC', sans-serif;
            font-size: 1.2em; /* 稍大字体 */
            font-weight: 500; /* 中等粗细 */
            line-height: 1.8; /* 行距 */
            margin: 5px 0;
            padding-left: 10px; /* 左边距 */
            text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); /* 文字阴影 */
        }
        /* 移动端调整 */
        @media (max-width: 768px) {
            .memory-text {
                padding: 10px;
            }
            .memory-text p {
                font-size: 1em; /* 移动端字体稍小 */
                padding-left: 5px; /* 移动端左边距稍小 */
            }
            .cate ul li a {
                font-size: 1.2em; /* 移动端年份字体稍小 */
            }
            .left-column {
                padding-right: 5px; /* 移动端减少右边距 */
            }
            .cate {
                margin-top: 2px; /* 移动端进一步减少上边距 */
            }
        }
    </style>
</head>
<body>
    <!-- 头部（移动端） -->
    <header class="header-mobile d-lg-none">
        <div class="logo">
            <a href="/">Record China</a>
        </div>
        <div class="menu">
            <button class="btn">☰</button>
        </div>
    </header>

    <!-- 主体 -->
    <main class="main">
        <div class="container-fluid">
            <div class="row">
                <!-- 左侧：标题 + 年份 -->
                <div class="col-md-3 left-column">
                    <div class="title">
                        <div class="inner">
                            <h2>记录历史 传播真相</h2>
                        </div>
                    </div>
                    <div class="intro">
                        <div class="inner">
                            <p>这是一个流动的纪念碑，碑身上交错刻下三条时间线：</p>
                            <p>主线：详细讲述1989年4月至6月发生在中国的一场规<br>模浩大的民主抗争运动。</p>
                            <p>背景线：八九运动推动了改变国际政治格局的苏东剧<br>变，因此记录了当代全球民主自由抗争史。</p>
                            <p>生命线：被大历史裹挟的个体命运的挣扎，从未因时<br>光流逝而消散。</p>
                            <p>请点击以下年份浏览具体事件信息：<br>
                               浏览大图时可用键盘左右箭头切换</p>
                            <a href="/" class="btn underline"></a>
                        </div>
                    </div>
                    <div class="cate">
                        <div class="inner">
                            <ul id="year-list">
                                <!-- 年份列表由 JS 动态生成 -->
                            </ul>
                        </div>
                    </div>
                </div>
                <!-- 中间：图片内容 -->
                <div class="col-md-6 content">
                    <!-- 纪念文字 -->
                    <div class="memory-text">
                        <p>光之记忆，历史永存 | Memory in Light, History Endures</p>
                        <p>自由之光，照亮历史 | Light of Freedom, Illuminates History</p>
                        <p>记忆不灭，历史不泯 | Memory Never Fades, History Never Dies</p>
                        <p>燃起希望，争取自由 | Ignite Hope, Fight for Freedom</p>
                        <p>和平永存，自由不熄 | Peace Endures, Freedom Shines</p>
                        <p>点亮烛光，铭记历史 | Light a Candle, Remember History</p>
                        <p>传递希望，为自由发声 | Pass the Light of Hope, Speak for Freedom</p>
                        <p>转发历史，唤醒自由 | Share History, Awaken Freedom</p>
                        <p>记录记忆，守护自由 | Record Memory, Defend Freedom</p>
                    </div>
                    <section class="board" id="image-board">
                        <div class="row">
                            <!-- 图片将在这里动态生成 -->
                        </div>
                        <div class="pagination" id="pagination">
                            <!-- 分页按钮将在这里生成 -->
                        </div>
                    </section>
                </div>
                <!-- 右侧：占位 -->
                <div class="col-md-3"></div>
            </div>
        </div>
    </main>

    <!-- 放大图片的 Modal -->
    <div class="modal fade" id="imageModal" tabindex="-1" aria-labelledby="imageModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-lg">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="imageModalLabel">图片预览</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <img src="" id="modalImage" style="width: 100%; height: auto;">
                    <!-- 纪念文字 -->
                    <div class="memory-text">
                        <p>光之记忆，历史永存 | Memory in Light, History Endures</p>
                        <p>自由之光，照亮历史 | Light of Freedom, Illuminates History</p>
                        <p>记忆不灭，历史不泯 | Memory Never Fades, History Never Dies</p>
                        <p>燃起希望，争取自由 | Ignite Hope, Fight for Freedom</p>
                        <p>和平永存，自由不熄 | Peace Endures, Freedom Shines</p>
                        <p>点亮烛光，铭记历史 | Light a Candle, Remember History</p>
                        <p>传递希望，为自由发声 | Pass the Light of Hope, Speak for Freedom</p>
                        <p>转发历史，唤醒自由 | Share History, Awaken Freedom</p>
                        <p>记录记忆，守护自由 | Record Memory, Defend Freedom</p>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- 页脚 -->
    <footer class="footer d-none d-lg-block">
        <div class="container-fluid d-flex justify-content-between">
            <div class="footer-left">
                 <a href="/about" class="btn underline" style="color: #ffffff; font-weight: bold;">关于我们</a>
            </div>
            <div class="footer-right">
                <button class="btn active">中</button>
                <button class="btn">EN</button>
            </div>
        </div>
    </footer>

    <!-- Bootstrap JS -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
    <!-- 自定义 JS -->
    <script>
        // 图片路径和完整文件名列表
        const basePath = "qsewr/dasf/erwq/reeqr/rwer/ewqr/tttt/";
        const imageFiles = [
            "29760928.png", "29761201.png", "29771201.png", "29780101-1.png", "29780101-2.png",
            "29780101-3.png", "29780101.png", "29780114.png", "29780213.png", "29780214.png",
            "29780301.png", "29780401.png", "29780402.png", "29780403.png", "29780415.png",
            "29780416.png", "2978041602.png", "2978041603.png", "2978041701.png", "29780418.png",
            "2978041802.png", "2978041901.png", "29780425.png", "29780426.png", "2978042602.png",
            "29780427.png", "2978042702.png", "2978042703.png", "29780428.png", "2978042901.png",
            "2978043001.png", "2978050301.png", "29780504.png", "29780507.png", "29780509.png",
            "29780520.png", "2978052001.png", "29780521.png", "29780522.png", "2978052202.png",
            "29780523.png", "29780525.png", "29780528.png", "2978052802.png", "29780529.png",
            "29780531.png", "2978053102.png", "29780602.png", "2978060202.png", "2978060203.png",
            "2978060301.png", "2978060302.png", "2978060401.png", "2978060402.png", "2978060403.png",
            "2978060404.png", "2978060405.png", "2978060406.png", "2978060407.png", "2978060408.png",
            "2978060409.png", "2978060410.png", "2978060411.png", "2978060412.png", "2978060501.png",
            "2978060502.png", "2978060503.png", "2978060601.png", "2978060602.png", "2978060701.png",
            "2978060702.png", "29780801.png", "29780802.png", "29780803.png", "29780804.png",
            "29780805.png", "29780806.png", "2978080602.png"
        ];

        // 年份映射
        const yearMap = {
            "2976": "1986",
            "2977": "1988",
            "2978": "1989"
        };

        // 按年份分组
        const images = {};
        imageFiles.forEach(file => {
            const yearCode = file.slice(0, 4);
            const displayYear = yearMap[yearCode];
            if (!images[displayYear]) images[displayYear] = [];
            images[displayYear].push(`${basePath}${file}`);
        });

        // 动态生成年份列表
        const yearList = document.getElementById('year-list');
        const imageBoard = document.getElementById('image-board').querySelector('.row');
        const pagination = document.getElementById('pagination');
        let currentYear = null;
        let currentPage = 1;
        const imagesPerPage = 2;
        let currentModalIndex = 0; // 当前 Modal 图片索引

        Object.keys(images).forEach(year => {
            const li = document.createElement('li');
            const a = document.createElement('a');
            a.href = '#' + year;
            a.textContent = year;
            a.onclick = () => {
                currentYear = year;
                currentPage = 1;
                showImages(year, currentPage);
            };
            li.appendChild(a);
            yearList.appendChild(li);
        });

        // 显示图片函数（带增强分页）
        function showImages(year, page) {
            const yearImages = images[year] || [];
            const totalImages = yearImages.length;
            const totalPages = Math.ceil(totalImages / imagesPerPage);
            const startIndex = (page - 1) * imagesPerPage;
            const endIndex = startIndex + imagesPerPage;
            const displayedImages = yearImages.slice(startIndex, endIndex);

            imageBoard.innerHTML = displayedImages.length > 0
                ? displayedImages.map((img, index) => `
                    <div class="col-md-6">
                        <img src="${img}" alt="${year} 时间线图片" data-bs-toggle="modal" data-bs-target="#imageModal" onclick="showModal('${img}', ${startIndex + index})">
                    </div>`).join('')
                : '<p>该年份暂无图片。</p>';

            // 更新分页按钮
            pagination.innerHTML = '';
            if (totalPages > 1) {
                const firstButton = document.createElement('button');
                firstButton.textContent = '第一页';
                firstButton.disabled = page === 1;
                firstButton.onclick = () => showImages(year, 1);
                pagination.appendChild(firstButton);

                const prevButton = document.createElement('button');
                prevButton.textContent = '上一页';
                prevButton.disabled = page === 1;
                prevButton.onclick = () => showImages(year, page - 1);
                pagination.appendChild(prevButton);

                const startPage = Math.max(1, page - 3);
                const endPage = Math.min(totalPages, page + 3);
                for (let i = startPage; i <= endPage; i++) {
                    const pageButton = document.createElement('button');
                    pageButton.textContent = i;
                    pageButton.className = i === page ? 'active' : '';
                    pageButton.onclick = () => showImages(year, i);
                    pagination.appendChild(pageButton);
                }

                const nextButton = document.createElement('button');
                nextButton.textContent = '下一页';
                nextButton.disabled = page === totalPages;
                nextButton.onclick = () => showImages(year, page + 1);
                pagination.appendChild(nextButton);

                const lastButton = document.createElement('button');
                lastButton.textContent = '最后一页';
                lastButton.disabled = page === totalPages;
                lastButton.onclick = () => showImages(year, totalPages);
                pagination.appendChild(lastButton);

                const jumpInput = document.createElement('input');
                jumpInput.type = 'number';
                jumpInput.min = 1;
                jumpInput.max = totalPages;
                jumpInput.placeholder = '页码';
                pagination.appendChild(jumpInput);

                const jumpButton = document.createElement('button');
                jumpButton.textContent = '跳转';
                jumpButton.onclick = () => {
                    const jumpPage = parseInt(jumpInput.value);
                    if (jumpPage >= 1 && jumpPage <= totalPages) showImages(year, jumpPage);
                };
                pagination.appendChild(jumpButton);
            }

            // 更新选中状态
            document.querySelectorAll('.cate a').forEach(a => a.classList.remove('active'));
            document.querySelector(`.cate a[href="#${year}"]`).classList.add('active');
        }

        // 显示放大图片的 Modal 并记录索引
        function showModal(imageSrc, index) {
            const modalImage = document.getElementById('modalImage');
            modalImage.src = imageSrc;
            currentModalIndex = index;
        }

        // 键盘翻页逻辑
        document.addEventListener('keydown', (event) => {
            if (!currentYear || !document.getElementById('imageModal').classList.contains('show')) return;

            const yearImages = images[currentYear];
            if (!yearImages || yearImages.length === 0) return;

            if (event.key === 'ArrowLeft') {
                currentModalIndex = (currentModalIndex - 1 + yearImages.length) % yearImages.length;
                document.getElementById('modalImage').src = yearImages[currentModalIndex];
            } else if (event.key === 'ArrowRight') {
                currentModalIndex = (currentModalIndex + 1) % yearImages.length;
                document.getElementById('modalImage').src = yearImages[currentModalIndex];
            }
        });

        // 默认显示第一个年份的图片
        if (Object.keys(images).length > 0) {
            currentYear = Object.keys(images)[0];
            showImages(currentYear, 1);
        }
    </script>
</body>
</html>
