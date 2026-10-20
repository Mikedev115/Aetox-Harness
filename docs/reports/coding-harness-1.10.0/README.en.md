# Aetox 1.10.0 on real coding tasks, same model: vs OpenCode · Codex CLI · omp · pi

> Results record · 11 tasks × 2 models × 3 runs = 66 runs per harness · 9 harness versions (Aetox 1.10.0 · 1.9.27 · 1.9.7 · 1.9.3 · 1.9.0 · OpenCode · Codex CLI · omp · pi) = 594, minus 1 excluded = **593 runs**, collected over several measurement rounds since 27 Sep 2026.
> New in this edition: 66 runs (Aetox 1.10.0). The rest come from earlier editions. The CWEval ablation and the GPT-6 Sol runs are in the [1.9.27 edition](../coding-harness-1.9.27) (Thai).
> [ภาษาไทย](README.md) · This is a translation of the Thai report; if the two disagree, the Thai one is the record.
>
> One question: **give the same instruction to the same model; which harness gets the model to ship work that passes more tests?**
> Quality of the work is the main measure. Time and tokens are shown as the price paid, not as the contest.
> Every task is scored by hidden tests. The harness never sees them while it works. Nobody grades by feel.
>
> Every cell is 3 runs. The large number is the mean; the small line under it is the best run and which run it was.
> Same task set as the [1.9.0 report](../coding-harness-2026-09-27). Other harnesses' results come from that measurement round.

---

## Where Aetox leads, and where it loses

**Read this first: runs vary a lot.** The 1.9.7 edition was measured twice on the same day, and SlopCodeBench moved 1.5–2.5 points between the two sets.
A gap of a few points in this report reads as "about the same". The conclusions below are gaps larger than that, or ones that repeat on every run.

**vs Aetox 1.9.27**

- **Answers now come in the language of the task.** English tasks: Luna answered in English 13 → 63 of 63 answers · Terra 47 → 63 of 63. Tasks written in Thai still got Thai answers 9 / 9 on both models (section 2).
- **Terra improved in almost every column, and got faster.** SlopCodeBench 93.8 → 95.2% · checkpoints fully passed 6.0 → 7.0 · Web-Bench 4.7 → 5.3 · CWEval func-sec@1 93.3 → 97.8% · Thai tasks 93.3 → 100% · time 70 → 63 min.
- **Luna got faster again and fixed the reported bug on every run.** Time 31 → 26 min · SWE-bench per bug report 1 → 3 of 3 runs · func-sec@1 91.1 → 93.3% · func@1 93.3 → 97.8%.
  SlopCodeBench 92.4 → 91.6% and Web-Bench 5.7 → 5.3 are within run-to-run noise (A1 run 2 and S3 run 3 were lower than the others).
- **Tokens and cost fell slightly.** Uncached input: Luna −6% · Terra −5%. Cost: Luna −2% · Terra −5% (section 3).

**vs other harnesses (same model)**

- **Secure code (CWEval func-sec@1).** Aetox leads on both models: Luna 93.3% vs 68.9% for the best other harness · Terra 97.8% vs 77.8%.
  **Almost all of this gap comes from the security skill that ships with Aetox.** With that skill turned off, Aetox lands in the same range as the others ([1.9.27 edition, section 2](../coding-harness-1.9.27#2-ชนะ-cweval-เพราะอะไร), Thai).
- **Multi-turn work (SlopCodeBench).** Aetox is highest in the table on both models, in both % tests passed and checkpoints fully passed: Terra 95.2% and 7.0 / 14 vs 92.8% and 6.7 for OpenCode · Luna 91.6% and 4.0 vs 90.8% and 3.7.
- **Time on Luna.** 26 min, the fastest in the table (pi 29 · Codex CLI 30).
- **Aetox + Luna vs other harnesses + Terra.** SlopCodeBench 91.6% sits in the range of the other harnesses on Terra (91.3–92.8%), in 26 min vs 46–63 min, on a model about 20× cheaper per token.

**Where Aetox loses**

- **Official SWE-bench on Luna.** 1 of 3 runs. The other two fixed `client_ip` as the bug report asked, but not `remote_ip`, which the dataset's tests also check. Other harnesses on Luna scored 33–67%.
- **Time on Terra.** 63 min, equal to OpenCode but slower than pi (46) and Codex CLI (54).
- **Tokens are still higher than the others.** Luna 6.4 M (others 0.8–4.5) · Terra 8.3 M (others 1.8–7.5). About 90% of that is cached input.

### Who Luna and Terra are

| | Input / output price (per 1 M tokens) | Cached input | Generation |
|---|---|---|---|
| GPT-6 Luna | $0.10 / $0.50 | $0.01 | GPT-6 |
| GPT-5.6 Terra | $2 / $12 | $0.20 | GPT-5.6 |

OpenAI list prices from the model catalogue (models.dev). Luna is the budget GPT-6 model, about 20× cheaper than Terra; Terra is the large GPT-5.6 model.
Aetox 1.10.0 cost per round of 11 tasks at these prices: **Luna $0.16 · Terra $4.99**.
Every harness in this report ran through a ChatGPT account, which is not billed per token, so dollar figures are for comparing scale only.

---

## 1. Scores (mean of 3 runs)

Each column uses its suite's own metric · **bold** = best in the column for that model · small line = best run · which run (for time and tokens, best = lowest).
The Aetox 1.9.27 row is the same data as in the 1.9.27 edition.

### GPT-6 Luna · reasoning: medium

| Harness | SlopCodeBench<br>% tests passed | SlopCodeBench<br>checkpoints fully passed | SWE-bench ML<br>resolved (1 task) | SWE-bench ML<br>per bug report † | Web-Bench<br>tasks passed | CWEval<br>func-sec@1 | CWEval<br>func@1 | Thai tasks \* | Time<br>(min) | Tokens<br>(M) |
|---|---|---|---|---|---|---|---|---|---|---|
| Aetox 1.10.0 | 91.6%<br><sub>best 92.7 · run 1</sub> | **4.0 / 14**<br><sub>best 4.0 · all runs</sub> | 33%<br><sub>best 100 · run 1</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | **93.3%**<br><sub>best 100.0 · run 1</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | **26**<br><sub>best 24 · run 2</sub> | 6.4<br><sub>best 6.3 · run 2</sub> |
| Aetox 1.9.27 | **92.4%**<br><sub>best 93.4 · run 2</sub> | 3.7 / 14<br><sub>best 4.0 · run 2, 3</sub> | 33%<br><sub>best 100 · run 1</sub> | 33%<br><sub>best 100 · run 1</sub> | **5.7 / 6**<br><sub>best 6.0 · run 1, 3</sub> | 91.1%<br><sub>best 93.3 · run 2, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 31<br><sub>best 28 · run 3</sub> | 6.1<br><sub>best 5.8 · run 3</sub> |
| OpenCode 1.18.32 | 90.8%<br><sub>best 93.8 · run 2</sub> | 3.7 / 14<br><sub>best 5.0 · run 2</sub> | **67%**<br><sub>best 100 · run 2, 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 4.3 / 6<br><sub>best 5.0 · run 2, 3</sub> | 66.7%<br><sub>best 66.7 · all runs</sub> | 86.7%<br><sub>best 86.7 · all runs</sub> | 94.1%<br><sub>best 97.8 · run 2, 3</sub> | 46<br><sub>best 39 · run 1</sub> | 3.4<br><sub>best 3.1 · run 2</sub> |
| Codex CLI 0.157.1 | 86.9%<br><sub>best 88.4 · run 2</sub> | 2.3 / 14<br><sub>best 3.0 · run 1</sub> | 33%<br><sub>best 100 · run 2</sub> | 67%<br><sub>best 100 · run 2, 3</sub> | 5.0 / 6<br><sub>best 6.0 · run 1</sub> | 66.7%<br><sub>best 66.7 · all runs</sub> | 88.9%<br><sub>best 93.3 · run 2</sub> | 97.8%<br><sub>best 100.0 · run 3</sub> | 30<br><sub>best 27 · run 1</sub> | 2.1<br><sub>best 2.0 · run 2</sub> |
| omp 18.3.4 | 88.7%<br><sub>best 91.1 · run 3</sub> | 2.3 / 14<br><sub>best 3.0 · run 2</sub> | 33%<br><sub>best 100 · run 1</sub> | 67%<br><sub>best 100 · run 1, 2</sub> | 4.5 / 6<br><sub>best 5.0 · run 3</sub> | 68.9%<br><sub>best 73.3 · run 3</sub> | 86.7%<br><sub>best 86.7 · all runs</sub> | 88.9%<br><sub>best 97.8 · run 2, 3</sub> | 45<br><sub>best 43 · run 2</sub> | 4.5<br><sub>best 4.5 · run 2</sub> |
| pi 0.87.1 | 87.8%<br><sub>best 90.6 · run 3</sub> | 2.3 / 14<br><sub>best 4.0 · run 3</sub> | 33%<br><sub>best 100 · run 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 3</sub> | 64.4%<br><sub>best 73.3 · run 3</sub> | 86.7%<br><sub>best 86.7 · all runs</sub> | 94.1%<br><sub>best 97.8 · run 1, 3</sub> | 29<br><sub>best 28 · run 2</sub> | **0.8**<br><sub>best 0.8 · run 1</sub> |

### GPT-5.6 Terra · reasoning: medium

| Harness | SlopCodeBench<br>% tests passed | SlopCodeBench<br>checkpoints fully passed | SWE-bench ML<br>resolved (1 task) | SWE-bench ML<br>per bug report † | Web-Bench<br>tasks passed | CWEval<br>func-sec@1 | CWEval<br>func@1 | Thai tasks \* | Time<br>(min) | Tokens<br>(M) |
|---|---|---|---|---|---|---|---|---|---|---|
| Aetox 1.10.0 | **95.2%**<br><sub>best 96.5 · run 1</sub> | **7.0 / 14**<br><sub>best 8.0 · run 3</sub> | 67%<br><sub>best 100 · run 2, 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 63<br><sub>best 55 · run 2</sub> | 8.3<br><sub>best 7.9 · run 2</sub> |
| Aetox 1.9.27 | 93.8%<br><sub>best 94.8 · run 3</sub> | 6.0 / 14<br><sub>best 7.0 · run 3</sub> | 67%<br><sub>best 100 · run 1, 2</sub> | **100%**<br><sub>best 100 · all runs</sub> | 4.7 / 6<br><sub>best 5.0 · run 1, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 70<br><sub>best 66 · run 2</sub> | 8.8<br><sub>best 8.1 · run 2</sub> |
| OpenCode 1.18.32 | 92.8%<br><sub>best 93.9 · run 3</sub> | 6.7 / 14<br><sub>best 7.0 · run 1, 2</sub> | 33%<br><sub>best 100 · run 1</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.0 / 6<br><sub>best 5.0 · all runs</sub> | 77.8%<br><sub>best 80.0 · run 1, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 97.8 · run 2</sub> | 63<br><sub>best 58 · run 1</sub> | 5.1<br><sub>best 5.0 · run 3</sub> |
| Codex CLI 0.157.1 | 91.5%<br><sub>best 92.5 · run 3</sub> | 5.3 / 14<br><sub>best 7.0 · run 1</sub> | 67%<br><sub>best 100 · run 1, 2</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.0 / 6<br><sub>best 5.0 · all runs</sub> | 73.3%<br><sub>best 80.0 · run 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 92.6%<br><sub>best 97.8 · run 3</sub> | 54<br><sub>best 49 · run 1</sub> | 3.8<br><sub>best 3.5 · run 1</sub> |
| omp 18.3.4 | 86.4%<br><sub>best 93.8 · run 2</sub> | 5.3 / 14<br><sub>best 6.0 · run 2, 3</sub> | 33%<br><sub>best 100 · run 1</sub> | **100%**<br><sub>best 100 · all runs</sub> | **5.7 / 6**<br><sub>best 6.0 · run 1, 2</sub> | 77.8%<br><sub>best 80.0 · run 1, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 91.1%<br><sub>best 91.1 · all runs</sub> | 71<br><sub>best 65 · run 3</sub> | 7.5<br><sub>best 6.8 · run 3</sub> |
| pi 0.87.1 | 91.3%<br><sub>best 92.8 · run 1</sub> | 6.0 / 14<br><sub>best 7.0 · run 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | **100%**<br><sub>best 100 · all runs</sub> | **5.7 / 6**<br><sub>best 6.0 · run 2, 3</sub> | 77.8%<br><sub>best 80.0 · run 1, 2</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 95.6%<br><sub>best 97.8 · run 1, 2</sub> | **46**<br><sub>best 44 · run 3</sub> | **1.8**<br><sub>best 1.6 · run 1</sub> |

| Metric | Meaning |
|---|---|
| SlopCodeBench % tests passed | Share of tests passed per task, averaged over 6 tasks (A1 B1 S1–S4). Every checkpoint reached is checked after the last turn, including earlier ones that may have broken. |
| SlopCodeBench checkpoints fully passed | Checkpoints where every test passes, out of 14 (A1 3 · B1 3 · S1–S4 2 each). |
| SWE-bench ML resolved | By the dataset's tests (FAIL_TO_PASS). One bug only (caddy-6350) × 3 runs, so values are 0 / 33 / 67 / 100%. Read it as a weak signal. |
| SWE-bench ML per bug report † | Same tests, with only the `remote_ip` case removed, leaving what the bug report actually asked for. |
| Web-Bench tasks passed | Tasks (task-N) whose Playwright specs all pass after all 6 tasks are done. |
| CWEval func-sec@1 | Functions that both work correctly and pass the attack tests, out of 15 · func@1 = works correctly only. |
| Thai tasks \* | Tests passed across F1 + F2 (45 tests). Written by the Aetox team, and Aetox ships a Thai-language skill. Read separately from the standard suites. |
| Time / tokens | Total for all 11 tasks per run · tokens = input (including cached) + output, as each harness reports them. |

### Aetox + Luna vs other harnesses on Terra

A smaller model on one harness against a larger model on the others.

| | Aetox 1.10.0 + Luna | OpenCode + Terra | Codex CLI + Terra | pi + Terra |
|---|---|---|---|---|
| SlopCodeBench % tests passed | 91.6%<br><sub>best 92.7 · run 1</sub> | 92.8%<br><sub>best 93.9 · run 3</sub> | 91.5%<br><sub>best 92.5 · run 3</sub> | 91.3%<br><sub>best 92.8 · run 1</sub> |
| SWE-bench ML resolved (1 task) | 33%<br><sub>best 100 · run 1</sub> | 33%<br><sub>best 100 · run 1</sub> | 67%<br><sub>best 100 · run 1, 2</sub> | 100%<br><sub>best 100 · all runs</sub> |
| SWE-bench ML per bug report † | 100%<br><sub>best 100 · all runs</sub> | 100%<br><sub>best 100 · all runs</sub> | 100%<br><sub>best 100 · all runs</sub> | 100%<br><sub>best 100 · all runs</sub> |
| Web-Bench tasks passed | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | 5.0 / 6<br><sub>best 5.0 · all runs</sub> | 5.0 / 6<br><sub>best 5.0 · all runs</sub> | 5.7 / 6<br><sub>best 6.0 · run 2, 3</sub> |
| CWEval func-sec@1 | 93.3%<br><sub>best 100.0 · run 1</sub> | 77.8%<br><sub>best 80.0 · run 1, 3</sub> | 73.3%<br><sub>best 80.0 · run 3</sub> | 77.8%<br><sub>best 80.0 · run 1, 2</sub> |
| SlopCodeBench checkpoints fully passed | 4.0 / 14<br><sub>best 4.0 · all runs</sub> | 6.7 / 14<br><sub>best 7.0 · run 1, 2</sub> | 5.3 / 14<br><sub>best 7.0 · run 1</sub> | 6.0 / 14<br><sub>best 7.0 · run 3</sub> |
| Time (min) | 26<br><sub>best 24 · run 2</sub> | 63<br><sub>best 58 · run 1</sub> | 54<br><sub>best 49 · run 1</sub> | 46<br><sub>best 44 · run 3</sub> |
| Tokens (M) | 6.4<br><sub>best 6.3 · run 2</sub> | 5.1<br><sub>best 5.0 · run 3</sub> | 3.8<br><sub>best 3.5 · run 1</sub> | 1.8<br><sub>best 1.6 · run 1</sub> |

---

## 2. What changed since 1.9.27

Every item is a change to the system, not more instructions for the model. Runs in this edition used a build from commit `f1e728d5`, which has all of the items below.

| Area | 1.9.27 | 1.10.0 |
|---|---|---|
| Reply language | The prompt told the model in three places to "use the user's language" without saying which language that was. | The engine detects the language from the user's own message and tells the model at the end of that message. If unsure, it says nothing. The three prompt lines were removed. |
| Results of writing or editing files | Reported only the file just written. | Also reports which files the agent has already changed in this turn, and by how many lines, so the model does not ask git again. |
| Skills that always need their companion files | The model opened a skill, then needed another call to read its companion files. | A skill can mark files it always needs; they arrive with the skill (the security skill uses this). |
| Repeated identical tool results | git and repo_map called again sent the whole result back again. | If the result is unchanged, only a pointer to the earlier identical result is sent. |

### Reply language (every answer in every turn)

| | Luna 1.9.27 | Luna 1.10.0 | Terra 1.9.27 | Terra 1.10.0 |
|---|---|---|---|---|
| English tasks answered in English | 13 / 63 | **63 / 63** | 47 / 63 | **63 / 63** |
| Answered in another language | Thai 47 · Arabic 2 · Hebrew 1 | — | Thai 16 | — |
| Thai tasks (F1 F2 SEC1) answered in Thai | 9 / 9 | 9 / 9 | 9 / 9 | 9 / 9 |

Counted from the letters of each answer, excluding code and file names · SEC1 was instructed in Thai, so it sits in the Thai row.

---

## 3. The price paid: why Aetox uses more tokens and time

Aetox is built to prove work before it says the work is done: run the project's tests or checks, write tests itself when the project has none, then fix until they pass.
Every check is another round that sends the whole conversation to the model again, which costs tokens and time.

| Harness | Check commands per run | Runs that wrote their own test file |
|---|---|---|
| Aetox 1.10.0 | 6.6 | 14% |
| Aetox 1.9.27 | 7.0 | 11% |
| Aetox 1.9.7 | 6.8 | 12% |
| Aetox 1.9.3 | **7.2** | **15%** |
| Aetox 1.9.0 | 7.0 | **15%** |
| OpenCode 1.18.32 | 6.7 | 9% |
| Codex CLI 0.157.1 | 5.1 | 2% |
| omp 18.3.4 | 5.0 | 3% |
| pi 0.87.1 | 3.8 | 5% |

Mean per run over 11 tasks × 2 models × 3 runs · same definitions as the 1.9.7 edition.

### Tokens split by cache (mean per round of 11 tasks)

Most input is earlier text the provider has cached (hit), which is priced several times lower than new input.
The numbers that reflect the real load are **uncached (miss)** and **output**. Each harness reports tokens by its own definitions, so compare them roughly.

| Harness | Luna input | Luna miss | Luna output | Terra input | Terra miss | Terra output |
|---|---|---|---|---|---|---|
| Aetox 1.10.0 | 6.30 M | **0.55 M** | 0.09 M | 8.17 M | **0.76 M** | 0.17 M |
| Aetox 1.9.27 | 5.98 M | 0.58 M | 0.09 M | 8.66 M | 0.81 M | 0.17 M |
| OpenCode 1.18.32 | 3.34 M | 0.46 M | 0.07 M | 4.92 M | 0.58 M | 0.13 M |
| Codex CLI 0.157.1 | 1.99 M | 0.33 M | 0.06 M | 3.69 M | 0.44 M | 0.13 M |
| omp 18.3.4 | 4.45 M | 0.50 M | 0.07 M | 7.40 M | 0.67 M | 0.14 M |
| pi 0.87.1 | 0.74 M | 0.25 M | 0.05 M | 1.66 M | 0.40 M | 0.10 M |

The full split, including cache-hit rates and Aetox 1.9.7, is in the [Thai report, section 3](README.md#3-ราคาที่จ่าย-ทำไม-aetox-ใช้-token-และเวลามากกว่า).

### What benchmark numbers measure, and what they do not

A benchmark gives one instruction and nobody answers questions. Every harness gets the same instruction, byte for byte, so the numbers mostly reflect the main agent working alone.
In larger real work, helper agents take reading and repetitive loops off the main agent, so it does not carry the whole context ([report](../LIVING-REALM-DELEGATION-TOKENS-20261004.md), Thai).
High token use in a benchmark is therefore not the whole picture of real-world cost.

---

## 4. From 1.9.0 to 1.10.0

### GPT-6 Luna · reasoning: medium

| Harness | SlopCodeBench<br>% tests passed | SlopCodeBench<br>checkpoints fully passed | SWE-bench ML<br>resolved (1 task) | SWE-bench ML<br>per bug report † | Web-Bench<br>tasks passed | CWEval<br>func-sec@1 | CWEval<br>func@1 | Thai tasks \* | Time<br>(min) | Tokens<br>(M) |
|---|---|---|---|---|---|---|---|---|---|---|
| Aetox 1.10.0 | 91.6%<br><sub>best 92.7 · run 1</sub> | **4.0 / 14**<br><sub>best 4.0 · all runs</sub> | **33%**<br><sub>best 100 · run 1</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | 93.3%<br><sub>best 100.0 · run 1</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | **26**<br><sub>best 24 · run 2</sub> | 6.4<br><sub>best 6.3 · run 2</sub> |
| Aetox 1.9.27 | **92.4%**<br><sub>best 93.4 · run 2</sub> | 3.7 / 14<br><sub>best 4.0 · run 2, 3</sub> | **33%**<br><sub>best 100 · run 1</sub> | 33%<br><sub>best 100 · run 1</sub> | **5.7 / 6**<br><sub>best 6.0 · run 1, 3</sub> | 91.1%<br><sub>best 93.3 · run 2, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 31<br><sub>best 28 · run 3</sub> | 6.1<br><sub>best 5.8 · run 3</sub> |
| Aetox 1.9.7 | 90.1%<br><sub>best 91.0 · run 2</sub> | 3.0 / 14<br><sub>best 4.0 · run 1</sub> | 0%<br><sub>best 0 · all runs</sub> | 33%<br><sub>best 100 · run 2</sub> | 5.0 / 6<br><sub>best 6.0 · run 2</sub> | 91.1%<br><sub>best 93.3 · run 2, 3</sub> | 95.6%<br><sub>best 100.0 · run 1</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 47<br><sub>best 46 · run 2</sub> | **4.8**<br><sub>best 4.5 · run 3</sub> |
| Aetox 1.9.3 | 91.8%<br><sub>best 93.3 · run 1</sub> | **4.0 / 14**<br><sub>best 5.0 · run 3</sub> | **33%**<br><sub>best 100 · run 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.0 / 6<br><sub>best 5.0 · all runs</sub> | **97.8%**<br><sub>best 100.0 · run 2, 3</sub> | **97.8%**<br><sub>best 100.0 · run 2, 3</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 47<br><sub>best 42 · run 2</sub> | 5.7<br><sub>best 5.2 · run 2</sub> |
| Aetox 1.9.0 | **92.4%**<br><sub>best 93.9 · run 3</sub> | 3.0 / 14<br><sub>best 4.0 · run 3</sub> | **33%**<br><sub>best 100 · run 1</sub> | 67%<br><sub>best 100 · run 1, 3</sub> | **5.7 / 6**<br><sub>best 6.0 · run 1, 3</sub> | 91.1%<br><sub>best 100.0 · run 3</sub> | **97.8%**<br><sub>best 100.0 · run 1, 3</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 64<br><sub>best 61 · run 3</sub> | 6.3<br><sub>best 5.8 · run 2</sub> |

### GPT-5.6 Terra · reasoning: medium

| Harness | SlopCodeBench<br>% tests passed | SlopCodeBench<br>checkpoints fully passed | SWE-bench ML<br>resolved (1 task) | SWE-bench ML<br>per bug report † | Web-Bench<br>tasks passed | CWEval<br>func-sec@1 | CWEval<br>func@1 | Thai tasks \* | Time<br>(min) | Tokens<br>(M) |
|---|---|---|---|---|---|---|---|---|---|---|
| Aetox 1.10.0 | **95.2%**<br><sub>best 96.5 · run 1</sub> | **7.0 / 14**<br><sub>best 8.0 · run 3</sub> | **67%**<br><sub>best 100 · run 2, 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **97.8%**<br><sub>best 100.0 · run 1, 2</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | **63**<br><sub>best 55 · run 2</sub> | 8.3<br><sub>best 7.9 · run 2</sub> |
| Aetox 1.9.27 | 93.8%<br><sub>best 94.8 · run 3</sub> | 6.0 / 14<br><sub>best 7.0 · run 3</sub> | **67%**<br><sub>best 100 · run 1, 2</sub> | **100%**<br><sub>best 100 · all runs</sub> | 4.7 / 6<br><sub>best 5.0 · run 1, 3</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 70<br><sub>best 66 · run 2</sub> | 8.8<br><sub>best 8.1 · run 2</sub> |
| Aetox 1.9.7 | 92.0%<br><sub>best 94.6 · run 3</sub> | **7.0 / 14**<br><sub>best 7.0 · all runs</sub> | 33%<br><sub>best 100 · run 2</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 1</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 95.6%<br><sub>best 100.0 · run 1</sub> | 76<br><sub>best 65 · run 2</sub> | **7.8**<br><sub>best 6.7 · run 2</sub> |
| Aetox 1.9.3 | 92.2%<br><sub>best 94.1 · run 1</sub> | 6.0 / 14<br><sub>best 7.0 · run 1</sub> | **67%**<br><sub>best 100 · run 2, 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | 5.3 / 6<br><sub>best 6.0 · run 3</sub> | 95.6%<br><sub>best 100.0 · run 3</sub> | 95.6%<br><sub>best 100.0 · run 3</sub> | **100.0%**<br><sub>best 100.0 · all runs</sub> | 81<br><sub>best 76 · run 1</sub> | 9.1<br><sub>best 8.8 · run 1</sub> |
| Aetox 1.9.0 | 94.3%<br><sub>best 94.8 · run 2</sub> | 6.3 / 14<br><sub>best 7.0 · run 1, 3</sub> | **67%**<br><sub>best 100 · run 1, 3</sub> | **100%**<br><sub>best 100 · all runs</sub> | **5.7 / 6**<br><sub>best 6.0 · run 1, 2</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 93.3%<br><sub>best 93.3 · all runs</sub> | 95.6%<br><sub>best 100.0 · run 2</sub> | 88<br><sub>best 78 · run 1</sub> | 9.4<br><sub>best 9.0 · run 2</sub> |

- **Terra is the best of every Aetox version on multi-turn work and security.** SlopCodeBench 94.3 → 92.2 → 92.0 → 93.8 → 95.2% (1.9.0 → 1.9.3 → 1.9.7 → 1.9.27 → 1.10.0) · func-sec@1 97.8%, higher than any earlier version.
- **Time improved on every version.** Luna 64 → 47 → 47 → 31 → 26 min · Terra 88 → 81 → 76 → 70 → 63 min.
- **Luna multi-turn work stays in the same range.** SlopCodeBench 90.1–92.4% on every version · checkpoints fully passed 4.0, equal to 1.9.3, the highest.
- **Tokens close to 1.9.27.** Luna 6.1 → 6.4 M · Terra 8.8 → 8.3 M, still more than 1.9.7 (4.8 · 7.8).

---

## 5. Per-task scores

Number = mean tests passed / total tests (multi-turn tasks count every turn) · in brackets: runs 1, 2, 3.
B3a counts runs that fixed the bug, out of 3: "official" uses the dataset's tests (both `client_ip` and `remote_ip` must be fixed) · "per bug report †" checks only `client_ip`, which the report asked for.

### GPT-6 Luna

| Task | Suite | Aetox 1.10.0 | Aetox 1.9.27 | OpenCode 1.18.32 | Codex CLI 0.157.1 | omp 18.3.4 | pi 0.87.1 |
|---|---|---|---|---|---|---|---|
| A1 | SlopCodeBench circuit_eval | 192.3 / 205 <sub>(198, 176, 203)</sub> | 203.7 / 205 <sub>(203, 205, 203)</sub> | 199.3 / 205 <sub>(199, 199, 200)</sub> | 198.7 / 205 <sub>(194, 202, 200)</sub> | 197.7 / 205 <sub>(188, 203, 202)</sub> | 191.0 / 205 <sub>(177, 197, 199)</sub> |
| B1 | SlopCodeBench database_migration | 78.7 / 87 <sub>(75, 76, 85)</sub> | 78.0 / 87 <sub>(76, 83, 75)</sub> | 75.7 / 87 <sub>(64, 84, 79)</sub> | 72.3 / 87 <sub>(68, 70, 79)</sub> | 68.3 / 87 <sub>(64, 63, 78)</sub> | 68.7 / 87 <sub>(66, 74, 66)</sub> |
| S1 | SlopCodeBench xjq | 48.7 / 51 <sub>(47, 50, 49)</sub> | 46.0 / 51 <sub>(48, 46, 44)</sub> | 48.7 / 51 <sub>(46, 51, 49)</sub> | 42.7 / 51 <sub>(37, 49, 42)</sub> | 46.7 / 51 <sub>(45, 48, 47)</sub> | 41.0 / 51 <sub>(29, 46, 48)</sub> |
| S2 | SlopCodeBench sheeteval | 37.0 / 44 <sub>(40, 35, 36)</sub> | 38.3 / 44 <sub>(37, 40, 38)</sub> | 35.3 / 44 <sub>(36, 37, 33)</sub> | 38.0 / 44 <sub>(33, 41, 40)</sub> | 36.0 / 44 <sub>(33, 38, 37)</sub> | 39.7 / 44 <sub>(41, 40, 38)</sub> |
| S3 | SlopCodeBench etl_pipeline | 64.7 / 73 <sub>(66, 69, 59)</sub> | 68.0 / 73 <sub>(70, 67, 67)</sub> | 66.0 / 73 <sub>(72, 62, 64)</sub> | 63.7 / 73 <sub>(67, 57, 67)</sub> | 67.0 / 73 <sub>(67, 68, 66)</sub> | 65.3 / 73 <sub>(67, 63, 66)</sub> |
| S4 | SlopCodeBench code_search | 24.3 / 25 <sub>(25, 25, 23)</sub> | 23.7 / 25 <sub>(23, 23, 25)</sub> | 23.7 / 25 <sub>(23, 25, 23)</sub> | 21.0 / 25 <sub>(23, 21, 19)</sub> | 23.0 / 25 <sub>(23, 23, 23)</sub> | 23.7 / 25 <sub>(23, 23, 25)</sub> |
| B3a | SWE-bench ML caddy-6350 | official 1/3<br>per bug report † 3/3 | official 1/3<br>per bug report † 1/3 | official 2/3<br>per bug report † 3/3 | official 1/3<br>per bug report † 2/3 | official 1/3<br>per bug report † 2/3 | official 1/3<br>per bug report † 3/3 |
| C1 | Web-Bench svg-solar | 22.3 / 23 <sub>(23, 22, 22)</sub> | 22.7 / 23 <sub>(23, 22, 23)</sub> | 20.7 / 23 <sub>(19, 21, 22)</sub> | 22.0 / 23 <sub>(23, 22, 21)</sub> | 20.5 / 23 <sub>(20, 21)</sub> | 21.7 / 23 <sub>(21, 21, 23)</sub> |
| SEC1 | CWEval Python | 14.0 / 15 <sub>(15, 13, 14)</sub> | 13.7 / 15 <sub>(13, 14, 14)</sub> | 10.0 / 15 <sub>(10, 10, 10)</sub> | 10.0 / 15 <sub>(10, 10, 10)</sub> | 10.3 / 15 <sub>(10, 10, 11)</sub> | 9.7 / 15 <sub>(9, 9, 11)</sub> |
| F1 | PromptPay QR \* | 28.0 / 28 <sub>(28, 28, 28)</sub> | 28.0 / 28 <sub>(28, 28, 28)</sub> | 26.3 / 28 <sub>(23, 28, 28)</sub> | 28.0 / 28 <sub>(28, 28, 28)</sub> | 24.0 / 28 <sub>(16, 28, 28)</sub> | 28.0 / 28 <sub>(28, 28, 28)</sub> |
| F2 | Thai tax invoice \* | 17.0 / 17 <sub>(17, 17, 17)</sub> | 17.0 / 17 <sub>(17, 17, 17)</sub> | 16.0 / 17 <sub>(16, 16, 16)</sub> | 16.0 / 17 <sub>(15, 16, 17)</sub> | 16.0 / 17 <sub>(16, 16, 16)</sub> | 14.3 / 17 <sub>(16, 11, 16)</sub> |

### GPT-5.6 Terra

| Task | Suite | Aetox 1.10.0 | Aetox 1.9.27 | OpenCode 1.18.32 | Codex CLI 0.157.1 | omp 18.3.4 | pi 0.87.1 |
|---|---|---|---|---|---|---|---|
| A1 | SlopCodeBench circuit_eval | 205.0 / 205 <sub>(205, 205, 205)</sub> | 204.7 / 205 <sub>(205, 204, 205)</sub> | 204.7 / 205 <sub>(205, 205, 204)</sub> | 196.0 / 205 <sub>(203, 192, 193)</sub> | 203.7 / 205 <sub>(202, 205, 204)</sub> | 201.0 / 205 <sub>(205, 205, 193)</sub> |
| B1 | SlopCodeBench database_migration | 80.7 / 87 <sub>(83, 78, 81)</sub> | 80.7 / 87 <sub>(80, 79, 83)</sub> | 79.7 / 87 <sub>(81, 76, 82)</sub> | 83.3 / 87 <sub>(83, 84, 83)</sub> | 77.0 / 87 <sub>(76, 75, 80)</sub> | 81.3 / 87 <sub>(84, 75, 85)</sub> |
| S1 | SlopCodeBench xjq | 47.0 / 51 <sub>(48, 47, 46)</sub> | 47.3 / 51 <sub>(47, 48, 47)</sub> | 48.7 / 51 <sub>(50, 50, 46)</sub> | 49.0 / 51 <sub>(50, 50, 47)</sub> | 32.3 / 51 <sub>(3, 48, 46)</sub> | 48.3 / 51 <sub>(47, 48, 50)</sub> |
| S2 | SlopCodeBench sheeteval | 40.7 / 44 <sub>(40, 41, 41)</sub> | 39.0 / 44 <sub>(38, 38, 41)</sub> | 38.7 / 44 <sub>(37, 38, 41)</sub> | 37.0 / 44 <sub>(37, 37, 37)</sub> | 39.0 / 44 <sub>(38, 41, 38)</sub> | 39.3 / 44 <sub>(41, 38, 39)</sub> |
| S3 | SlopCodeBench etl_pipeline | 72.3 / 73 <sub>(72, 72, 73)</sub> | 70.7 / 73 <sub>(71, 71, 70)</sub> | 69.7 / 73 <sub>(70, 70, 69)</sub> | 70.3 / 73 <sub>(70, 70, 71)</sub> | 71.0 / 73 <sub>(72, 71, 70)</sub> | 70.0 / 73 <sub>(72, 68, 70)</sub> |
| S4 | SlopCodeBench code_search | 23.7 / 25 <sub>(25, 23, 23)</sub> | 23.0 / 25 <sub>(23, 23, 23)</sub> | 21.7 / 25 <sub>(19, 23, 23)</sub> | 20.3 / 25 <sub>(19, 19, 23)</sub> | 20.3 / 25 <sub>(19, 23, 19)</sub> | 19.0 / 25 <sub>(19, 19, 19)</sub> |
| B3a | SWE-bench ML caddy-6350 | official 2/3<br>per bug report † 3/3 | official 2/3<br>per bug report † 3/3 | official 1/3<br>per bug report † 3/3 | official 2/3<br>per bug report † 3/3 | official 1/3<br>per bug report † 3/3 | official 3/3<br>per bug report † 3/3 |
| C1 | Web-Bench svg-solar | 22.0 / 23 <sub>(23, 22, 21)</sub> | 21.7 / 23 <sub>(22, 21, 22)</sub> | 21.7 / 23 <sub>(22, 21, 22)</sub> | 21.7 / 23 <sub>(22, 21, 22)</sub> | 22.7 / 23 <sub>(23, 23, 22)</sub> | 22.7 / 23 <sub>(22, 23, 23)</sub> |
| SEC1 | CWEval Python | 14.7 / 15 <sub>(15, 15, 14)</sub> | 14.0 / 15 <sub>(14, 14, 14)</sub> | 11.7 / 15 <sub>(12, 11, 12)</sub> | 11.0 / 15 <sub>(11, 10, 12)</sub> | 11.7 / 15 <sub>(12, 11, 12)</sub> | 11.7 / 15 <sub>(12, 12, 11)</sub> |
| F1 | PromptPay QR \* | 28.0 / 28 <sub>(28, 28, 28)</sub> | 28.0 / 28 <sub>(28, 28, 28)</sub> | 28.0 / 28 <sub>(28, 28, 28)</sub> | 27.0 / 28 <sub>(25, 28, 28)</sub> | 27.0 / 28 <sub>(25, 28, 28)</sub> | 27.0 / 28 <sub>(28, 28, 25)</sub> |
| F2 | Thai tax invoice \* | 17.0 / 17 <sub>(17, 17, 17)</sub> | 14.0 / 17 <sub>(14, 14, 14)</sub> | 14.0 / 17 <sub>(13, 16, 13)</sub> | 14.7 / 17 <sub>(16, 12, 16)</sub> | 14.0 / 17 <sub>(16, 13, 13)</sub> | 16.0 / 17 <sub>(16, 16, 16)</sub> |

\* F1 and F2 were written by the Aetox team (Thai-specific work with no standard suite), and Aetox ships a Thai-language skill.

Every run (time, tokens, tool calls) is in [results.csv](results.csv) · rows with `version` = `aetox 1.10.0` are this edition; other rows come from the 1.9.27 edition.

---

## 6. Measurement rules

**Excluded and re-run runs, same rules as earlier editions.** Every original run is kept in the raw results; nothing is deleted.

- Failed because of the environment (network drop, login collision, the machine closing the program mid-run) → excluded and re-run.
- A run scoring below 60% of the median of the other two runs in the same cell → re-run once, and the re-run counts whether it is better or not.
- A turn that hits the 30-minute ceiling does not count.
- A cell that scores 0 or fails every run: the failing tests are checked against the task before the number is accepted.

This edition had no failed runs, no re-runs, and no run below 60% of the median of the other two.

| Item | Setting |
|---|---|
| Models | `gpt-6-luna` and `gpt-5.6-terra`, reasoning medium, through a ChatGPT (Codex) account, the same account for every harness |
| Aetox 1.10.0 | CLI and engine built from commit `f1e728d5` on main · driven with `aetox chat`, one message at a time · full approval · a fresh data folder for every run · no user skills or settings |
| Aetox 1.9.27 · 1.9.7 · 1.9.3 · 1.9.0 | Results from the [1.9.27 edition](../coding-harness-1.9.27), same settings |
| OpenCode | 1.18.32 · `opencode run --auto` · a separate config folder holding only the login file |
| Codex CLI | 0.157.1 · `codex exec --ephemeral` with full approval · a separate `CODEX_HOME` holding only the login file; no `AGENTS.md`, plugins or user settings; web search off |
| omp | 18.3.4 · `omp -p --mode json --auto-approve` · empty home folder, logged in with the same ChatGPT account |
| pi | 0.87.1 · `pi -p --mode json` · empty home folder with only a copy of the login file · telemetry off |
| Messages sent | The same files, byte for byte, for every harness · multi-turn tasks are sent one turn at a time, nobody answers questions mid-run |
| Time limit | 30 minutes per turn |
| Runs | 3 per cell on one Windows 11 machine · 1.10.0 ran 5 runs in parallel |

**Fairness notes**

- Every harness used the same model, account, reasoning level, messages and starting files. Only the harness differs.
- Aetox ran with the skills its installer ships, as users get it. OpenCode, omp and pi ship no skills. Codex had its system skills, plugins and web search turned off, leaving only the harness and the login.
- Aetox versions were measured on later days than the other harnesses. Same model and account, but the provider side can change between days.
- Times were measured while several runs ran at once. They compare fairly with each other, but a single run on its own would be faster.

---

## 7. Reproduce it

Everything a harness sees, and everything that scores it, is published in [`../coding-harness-2026-09-27/tasks/`](../coding-harness-2026-09-27/tasks), with each upstream suite's licence.
Each task folder holds:

| Folder | What it is |
|---|---|
| `prompts/` | The exact messages sent, one file per turn |
| `start/` | The starting files. Every run begins from a fresh git repo containing only these |
| `hidden-tests/` | The tests used for scoring, copied in only after the harness has finished its turn |

| Task | Suite | Upstream task |
|---|---|---|
| A1 | SlopCodeBench | `circuit_eval` |
| B1 | SlopCodeBench | `database_migration` |
| S1–S4 | SlopCodeBench | `xjq` · `sheeteval` · `etl_pipeline` · `code_search` |
| B3a | SWE-bench Multilingual | `caddyserver__caddy-6350` |
| C1 | Web-Bench | `svg-solar` |
| SEC1 | CWEval | Python subset, 15 functions (`map.json` lists them) |
| F1 · F2 | Written by the Aetox team | PromptPay QR · Thai tax invoice |

To repeat a run with any harness:

1. Copy `tasks/<task>/start/` into an empty folder and `git init` it.
2. Send each file in `prompts/` to the harness, in order, one turn at a time, with full approval and a 30-minute limit per turn.
3. Copy `hidden-tests/` in and run them. For SlopCodeBench, install the run's own `requirements.txt` first, as the task instructs.
4. Compare with the matching rows in [results.csv](results.csv) (`task, model, harness, run, passed, tests, seconds, tokens_in, tokens_cached, tokens_out, tool_calls, b3a_issue_resolved, version, variant`).

The queue script that drove these runs is part of the Aetox source and is not published. The harness commands it ran are the ones in the table in section 6.
Aetox 1.10.0 is available as a [release](https://github.com/Mikedev115/Aetox-Harness/releases/tag/v1.10.0); the measured build was commit `f1e728d5`, before the release was tagged.
