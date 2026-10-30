<p align="center">
  <img src="docs/assets/logo.png" alt="Aetox" width="72">
</p>

<h1 align="center">Aetox harness</h1>

<h3 align="center">让 AI 作为一个系统工作的环境</h3>

<p align="center">
Aetox Harness 是运行在你电脑上的 AI 工作环境 —— 把模型、技能、团队、代码和记忆整合成一个系统。你用得越多，
它越懂你的工作方式。交回来的是一份可以直接打开使用的文件，在你自己的电脑上，用你已有的模型。
</p>

<table>
<tr>
<td width="50%" valign="top">

**助手**

### 融入你的工作，而不是让你去适应它

> *"两小时的会议录音 —— 帮我总结，并列出谁领了什么任务"*

交回来的是：带**议题**、**决议**和**每句话在第几分钟**的摘要，加上待办事项和负责人 —— 是你文件夹里的一个文件，
不是聊天框里的文字。

它分部分地认识你 —— 你的档案、这台电脑、每个工作台、每个 worker，以及你重复的做法 —— 因此能触及比通用助手
更深的上下文，而不必把全部内容塞进每一次请求。

`录制会议` · `收据` · `整理文件` · `磁盘空间` · `幻灯片` · `实时检索` · `浏览器` · `自动化` · `视频与内容` · `团队` · `待办与提醒`

<img src="docs/assets/readme-assistant-meeting.png" alt="一段会议录音的摘要：议题与决议，标注每句话在第几分钟，以及两条带负责人的待办" width="380">

</td>
<td width="50%" valign="top">

**开发者**

### 写出能跑的代码 —— 并记住这个项目定过什么

它在真实项目里工作：真实文件、编辑器、终端、语言服务器和 git。它会读项目里已有的 `AETOX.md`、`AGENTS.md`
或 `CLAUDE.md`，在说"完成"之前先跑项目自己的测试，并如实写下哪些部分它没有验证。

**记忆按项目分开** —— 打开哪个文件夹，就得到那个文件夹的记忆：你的编码习惯、已经定下的决策、团队的约定。
不用每次重讲，而聊天里提到相关的事时会自动取用。

**是系统让模型写出更好的代码** —— 这是我们唯一卖的东西，不是模型本身有多聪明。模型是你的，哪一个都行。

**别信我们 —— 自己跑一遍。** 所有题目、起始文件和隐藏测试都是公开的。
[自己跑任何一题](docs/reports/coding-harness-1.10.0/README.en.md) · [全部数字，包括速度和 token →](#测评结果)

`真实文件` · `终端` · `git` · `diagnostics` · `symbol` · `MCP 工具` · `项目自己的测试` · `子代理`

<img src="docs/assets/readme-code-20261013.jpg" alt="终端里 npm test 跑到通过，旁边是改过的文件与增删的行数" width="380">

</td>
</tr>
</table>

<p align="center">
  <sub>每一个数字都是同模型对比，题目、起始文件和隐藏测试全部公开，你可以自己跑。
  <a href="#测评结果">如何测量，以及全部数字</a></sub>
</p>

<p align="center">
  <a href="https://apps.microsoft.com/detail/9N4KKBRRSCZZ"><img src="docs/assets/microsoft-store-en.svg" alt="从 Microsoft Store 获取" width="256"></a>
</p>

<p align="center">
  <strong><a href="https://github.com/Mikedev115/Aetox/releases/latest/download/aetox-amd64-installer.exe">下载 Windows 安装程序（.exe）</a></strong>
  · <a href="https://github.com/Mikedev115/Aetox/releases/latest/download/aetox-cli-setup.exe">Aetox CLI</a>
  · <a href="#安装">Scoop 与便携版</a><br>
  Windows 10 及以上 · x64 · 免费使用 · 用你已有的模型，或在本地免费跑一个
</p>

<p align="center">
  <a href="https://mikedev115.github.io/Aetox-landing/">网站</a> ·
  <a href="docs/release-notes/">更新内容</a> ·
  <a href="docs/reports/README.md">测评报告</a> ·
  <a href="README.md">English</a> ·
  <a href="README.th.md">ภาษาไทย</a>
</p>

<p align="center">
  <a href="https://mikedev115.github.io/Aetox-landing/assets/hero-en.mp4"><img src="docs/assets/readme-hero-en.gif" alt="一个建站任务在 Aetox 里变成计划、工具步骤和做好的页面" width="1008"></a>
</p>

*真实应用的操作录像，不是测评跑分。*

## 开始使用

1. **安装** 从 Microsoft Store 或上面的安装程序。只选一条渠道并一直用它：Store 版和安装程序版各自保存设置、对话和密钥。
2. **接上模型** 打开 **设置 › 模型**，用 ChatGPT、GitHub Copilot 或 OpenRouter 登录，或粘贴 key，
   或指向你机器上的 LM Studio 或 Ollama。
3. **交一件事** 选 **助手** 处理文件、文档和网页，或选 **代码** 打开一个项目文件夹。

## 一条指令，步骤它自己安排

卖点不是"能做很多事"，而是：**你交出一件事，它自己判断先做什么后做什么**，每一步都用真实工具，
每一步的结果先检查再往下走 —— 每一步你都看得见，也随时可以叫停。

| 步骤 | 它自己做的事 |
|:---|:---|
| **1. 先读** | 打开项目，读 README、数据库结构、项目的运行方式，以及相关的文档 |
| **2. 计划** | 把事情拆成条目，说明哪一项要等哪一项，以及还有哪些不知道 |
| **3. 逐步动手** | 改文件、写 migration、跑命令、调用你接好的 MCP server、驱动真实浏览器 |
| **4. 检查自己的工作** | 跑测试、读语言服务器报的错、打开刚建好的页面自己点一遍 |
| **5. 如实汇报** | 哪些完成、哪些没通过、哪些它检查不了 |

**真的要用好几步的活**

- *"给我的店做一个预约系统"* —— 把需求读成规格 → 设计数据库 → 写 migration → 建页面 → 跑测试 →
  打开浏览器自己预约一次 → 汇报还差什么
- *"查五家竞品，做一张对比表"* —— 驱动浏览器开多个标签页，一页一页读 → 汇总 → 写成 `.xlsx`
- *"两小时的会议录音，帮我总结"* —— 转写 → 拆出议题和决议并标出第几分钟 → 取出待办和负责人 → 写成文件
- *"修这个 bug"* —— 读 stack trace → 找到相关文件 → 修 → 跑项目自己的测试 → 不过就再绕回来

**一件事一个人做不完** —— 交给团队，回来的是整套完成的成果，并说明谁做了什么。

## 为什么同一个模型做得更好

模型负责思考，其余由环境完成 —— 而且用的是真东西，不是模拟。

- **手** 自带标签页和配置的真实浏览器、真实终端、真正写出 `.docx` `.xlsx` `.pptx` 的工具，以及真实的剪辑室。
  它不会向你描述表格 —— 它把文件写出来，再打开看
- **眼** OCR 和文档读取器直接读扫描页，而不是让视觉模型去猜；读的是当时截下来的页面，而不是凭记忆
- **模型不可能知道的事** 技能里存着：幻灯片 1280×720 的画幅、攻击者怎么绕过路径检查、泰国发票怎么进位 ——
  以及你的工作方式
- **交回之前先检查** 跑测试、打开自己做出来的文件看一遍，并如实说出哪些它证明不了

## 它会学会你的工作方式

每天早上都从零开始的环境，是一个每天都要重新租的工具。Aetox 会把它了解到的你留下来。

- **只有重复出现才算一种做法** 一天只是一次事件；不同的三天才是你的工作方式。它不会凭一次对话就断定你是这样的人
- **由你说"是"** 任何在对话之外提出的内容都会等你，而它留下的每一行都可以从添加它的那条回答下面撤销
- **按"谁需要读"分开 —— 但真正干活时都能取到** 架构决定谁读什么：你的档案跟着每一项工作，关于这台电脑的跟着
  需要它的工作台，每个房间的做法留在那个房间，每个 worker 的留在那个 worker —— 但干活时它能检索全部历史、
  打开任何一天的笔记、取用你电脑里的文档，所以不必把一切塞进每次请求，也不会忘掉
- **它是一个普通 `.md` 文件文件夹，不是黑匣子** 你的档案、关于这台电脑的记录、每个房间的笔记、每个工作日一份的
  笔记，以及它看见你重复的路径 —— 都写在你能打开、修改、拷给别的 agent 或直接删掉的地方。它学到的东西没有一样
  被锁在只有 Aetox 能读的格式里

在哪里运行由你决定，模型也是你的：31 个 provider，云端和本地，模型列表在运行时向 provider 询问，
所以新模型不用等 Aetox 发新版。今天全部都在你的电脑上跑。
[全部 provider](https://mikedev115.github.io/Aetox-landing/features/#models)

## 测评结果

**代码** —— 五套公开题目 × 2 个 OpenAI 公开模型
（[`gpt-6-luna`](https://developers.openai.com/api/docs/models/gpt-6-luna)、
[`gpt-5.6-terra`](https://developers.openai.com/api/docs/models/gpt-5.6-terra)，medium 推理）× 3 轮，
对比另外四款 harness（OpenCode · Codex CLI · omp · pi）。每一题都由隐藏测试判定，同一个账号，
四款都不做任何用户设置。

| 题目 | 测什么 | Aetox | 其他最好的一款 |
|:---|:---|--:|--:|
| **CWEval** `func-sec@1` | 既功能正确、又通过攻击测试的函数（共 15 个） | **93.3%** (Luna)<br>**97.8%** (Terra) | omp 68.9%<br>OpenCode · omp · pi 77.8% |
| **SlopCodeBench** % 测试通过 | 六个多轮任务，检查旧代码是否被弄坏 | **91.6%** (Luna)<br>**95.2%** (Terra) | OpenCode 90.8%<br>OpenCode 92.8% |
| **SlopCodeBench** checkpoint 全通过 | 所有测试都通过的 checkpoint，共 14 个 | **4.0 / 14**<br>**7.0 / 14** | OpenCode 3.7<br>OpenCode 6.7 |
| **SWE-bench Multilingual**（caddy-6350） | 公开仓库里的真实 bug，官方判定标准 | 1/3 轮 (Luna)<br>2/3 轮 (Terra) | OpenCode 2/3<br>pi 3/3 |
| **Web-Bench** | 做完 6 个任务后，Playwright 规格完全通过的任务 | 5.3 / 6<br>5.3 / 6 | pi 5.3<br>omp · pi 5.7 |

CWEval 的领先来自随装即开的安全技能 —— 关掉它，Aetox 就落回其他几款同一区间。
而 SWE-bench 是我们没有赢的那一套，它就在这张表里。

**文档与媒体** —— 在同一个应用上，开启技能对比关闭技能

- **泰语幻灯片**：**98 对比 49**（关闭技能时）
- **一段原片**：17.5 秒 → **11.1 秒**，合并四段静音、切掉一处口头赘词，之后仍可继续剪
- **团队**：主代理的上下文 **107k 而不是 190k** ·
  [整个系统一起测](docs/reports/LIVING-REALM-DELEGATION-TOKENS-20261004.md)
- **它学会你**：三天里代理提出的 16 条关于主人的内容，留下 4 条，四条都是关于这个人；关于"机器"的 8 条一条都没过 ——
  这是标准在起作用，不是在宣传功能

这些全部公开，包括 prompt 和测试：
[最新的代码报告](docs/reports/coding-harness-1.10.0/README.en.md) ·
[技能开与关](docs/reports/SKILL-BENCH.md) · [全部报告](docs/reports/README.md)

## 怎么用才发挥得出来

Aetox 没有让模型变聪明的魔法 —— 它按架构把模型的能力榨出来。小模型也能撑起大系统，原因就在这里；
**所以为工作挑对模型同样重要**。

我们的上下文架构不是从「怎么省 token」出发，而是从「这一轮模型真正需要知道什么」出发。
系统知道的东西会被映射成自己的结构 —— 分成类别，作为节点彼此相连 —— 运行时只发送这项活真正用到的地方。
prompt 按每一项活重新组装，不是一开始就整块发出，也不会处理这项活用不到的部分：与请求无关时，你的个人资料不会跟着走。
记忆、技能和 MCP 全部建立在同一套架构之下。

技能正文和工具 schema 不会一次发出。模型看到的是一个只有名字的架子，真打开时才为正文付费，
所以五十个技能和一个技能同价；增加一个 MCP 服务器增加的是名字，不是整套 schema。
而且每个 agent 都可以各自选择要接的技能和 MCP。

需要看到全局的活，我们不让模型一片一片地要，而是从头一次把广度交出去，活在一次对话里做完。
token 就是花在这里：总输入是其他 harness 的 1.1–8.5 倍，但其中约 **91% 是缓存输入**，价格低好几倍。
**真正的负担是未缓存输入加输出，差距只有 1.1–2.1 倍**：Luna 0.64M 对比 0.30–0.57M · Terra 0.93M 对比 0.50–0.81M。

所以在某些代码任务上多花 token 不是浪费，而是把模型能力榨出来、换回更好结果的副作用，
普通任务上并不发生 —— 按指令加一个 MCP 或一次装一批技能，全部由上面这套机制处理，实际用起来便宜得出乎意料。

**这是值得做的交换**，不是该害怕的事。真想降低，就按工作挑模型。

- **要读整个项目、或连续跑很久的活** → 大模型占优
- **重复做的日常活** → 小模型够用，系统会补上差的部分
- **一切都跑在 Windows 上** —— 屏幕、浏览器和电脑控制层全是 Win32

<sub>每轮的分数摆动 1.5–2.5 个点，差几个点就当作打平 · 测量方法和每一轮结果都在
[报告里](docs/reports/coding-harness-1.10.0/README.en.md)</sub>

## 盒子里有什么

| | 做什么 | 更多 |
|:---|:---|:---|
| **开发者** | Agent 模式替你推进整个项目；Editor 模式让你自己写、它在旁边帮。语言服务器、Git、真实终端 | [代码](https://mikedev115.github.io/Aetox-landing/features/#code) |
| **助手** | 会议、收据、资料检索和原始数据，变成能打开的文件：摘要、Excel、Word、PowerPoint | [助手](https://mikedev115.github.io/Aetox-landing/features/#assistant) |
| **团队** | 把整件事交给现成的团队，看见谁做了什么、交回了什么、花了多少 | [团队](https://mikedev115.github.io/Aetox-landing/features/#teams) |
| **剪辑室** | 长视频变短片，带字幕，之后还能继续剪 | [剪辑室](https://mikedev115.github.io/Aetox-landing/features/#studio) |
| **手机** | 在 Android 或 iPhone 上用自己的 PC 上的 Aetox；密钥不离开 PC | [手机设置](https://mikedev115.github.io/Aetox-landing/download/#mobile-preview) |
| **CLI** | 终端里的「代码」工作台，与应用共用密钥、对话和记忆 | [安装](#安装) |

界面支持**英文、泰文和中文**。

## 模型

31 个 provider，云端和本地。ChatGPT（Codex）、GitHub Copilot 和 OpenRouter 可以直接登录；其余的用 API key
或本地服务器地址。LM Studio 和 Ollama 拿到和前沿模型一样的一套工具。对话中途换模型，整个上下文跟着走。

图片也可以留在你的机器上：把 Aetox 指向 ComfyUI 或 Stable Diffusion WebUI（Forge Neo、Forge 或 A1111），
用自己的 GPU 出图，包括 FLUX.2 klein。[本地出图](docs/LOCAL-IMAGES.md)

> **免费使用，但不含免费云端推理。** 真正干活需要一个模型：你已有的账号或 key，或者本地模型。
> 价格按你所选的 provider 计。

## 安装

| 渠道 | 安装 | 更新 |
|:---|:---|:---|
| **Microsoft Store** | [Store 页面](https://apps.microsoft.com/detail/9N4KKBRRSCZZ) 或 `winget install --id=9N4KKBRRSCZZ --source=msstore` | Windows 负责更新 |
| **安装程序** | [aetox-amd64-installer.exe](https://github.com/Mikedev115/Aetox/releases/latest/download/aetox-amd64-installer.exe) | 应用内：**下载更新 → 重启以更新** |
| **Aetox CLI** | [aetox-cli-setup.exe](https://github.com/Mikedev115/Aetox/releases/latest/download/aetox-cli-setup.exe) | 装在用户目录，不需要管理员权限，自动加入 PATH |
| **Scoop** | `scoop install https://raw.githubusercontent.com/Mikedev115/Aetox/main/scoop/aetox.json` | `scoop update aetox` |

安装程序还没有代码签名，所以 Windows 可能提示 **"Windows protected your PC"**：选 **More info → Run anyway**，
或使用 Microsoft 签名的 Store 版。便携版 zip、杀毒软件相关说明，以及更新没生效时怎么办：
[更新指南](docs/UPDATING.md)。

## 安全与你的数据

- **数据留在你的电脑上** —— 对话、文件、浏览器数据，以及它学到的一切都在你的磁盘上；中间没有我们的服务器，
  也不做统计。云端 provider 看到的，就是它的 API 本来会看到的内容。
- **密钥** 绑定你的 Windows 账号，并从日志和工具结果里剔除。
- **能做多少由你决定** —— 每次都问 · 只问有风险的 · 完全放开。对外发布内容和按名字结束程序，任何档位都会问。
- **绝不打开机密文件** —— `.ssh`、`.aws`、Windows 凭据存储、浏览器配置，以及 Aetox 自己的密钥文件，
  任何工具都打不开。
- **靠什么保证** —— 这些是 Aetox 自己代码里的检查，不是操作系统的沙箱。如果机器上有不能丢的数据，
  给 Aetox 单独的 API key，并放在虚拟机或 Windows Sandbox 里跑。

[安全细节](https://mikedev115.github.io/Aetox-landing/safety/) · 发现安全问题？[SECURITY.md](SECURITY.md) 说明如何私下报告

## 源码、许可与作者

Aetox 免费使用，但不是开源软件。可以装在任何数量的机器上，也可以用于商业工作；请不要修改、再分发、
换品牌或出售（[LICENSE](LICENSE)）。你写的技能、agent、prompt 和 MCP server 属于你，允许出售（第 4 条）。
本仓库的源码停在 v1.7.0，留在这里供阅读和审计。之后的版本只发二进制，所以你运行的版本没有源码可读；
[没有源码时你能自己核查什么](#安全与你的数据)。每个版本都在 `checksums.txt` 里列出所有文件的 SHA-256，
用应用的更新密钥签名，2026 年 10 月 21 日之后构建的版本还带 `sbom.txt`：编进程序的每个 Go module 和 web 包。
Go 部分可以用 `go version -m aetox.exe` 直接和二进制对照。Microsoft Store 版由 Microsoft 签名。

Aetox 由一个人写成。问题和想法：[Facebook 群组](https://www.facebook.com/share/g/1BnXC5EiWg/)。
Bug：[issues](https://github.com/Mikedev115/Aetox/issues) —— 应用里的 **设置** 会自动带上你的版本号，
以及已经去掉机密内容的日志。❤️ [支持这个项目](SPONSOR.md)

---

<p align="center">
  © 2026 Chayaphon Phromsawana · All rights reserved · <a href="LICENSE">Licence</a> · <a href="THIRD-PARTY-NOTICES.md">Third-party notices</a>
</p>
