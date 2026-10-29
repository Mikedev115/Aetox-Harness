// Names are presentation only. Availability always comes from the CLI.
const names: Record<string, string> = {
  'devin-swe-2-preview': 'SWE-2', 'devin-swe-1-7': 'SWE-1.7',
  'devin-swe-1-7-lightning': 'SWE-1.7 Lightning', 'devin-swe-1-7-paid': 'SWE-1.7 Priority',
  'fusion-fast': 'Fusion Fast', 'devin-research-preview': 'Research',
}
export function devinFamily(id: string) {
  return /^(swe-2(?:-lightning|-priority)?)-(medium|high|max)$/.exec(id)
}
export function devinLabel(id: string, effort = ''): string {
  if (effort && (!/-(medium|high|max|low|none)$/.test(id) || id.startsWith('fusion-'))) return `${devinLabel(id)} · ${title(effort)}`
  // ACP thought_level overrides the tier carried in the base model ID.
  if (effort && !id.startsWith('fusion-')) id = id.replace(/-(medium|high|max|low|none)$/, `-${effort}`)
  if (names[id]) return names[id]
  const family = devinFamily(id)
  if (family) return `SWE-2 ${title(family[2])}${family[1].includes('lightning') ? ' (Lightning)' : family[1].includes('priority') ? ' (Priority)' : ''}`
  return id.replace(/\d+(?:-\d+)+/g, version => version.replaceAll('-', '.')).split('-').map(title).join(' ')
}
function title(word: string): string {
  if (['gpt', 'glm', 'swe'].includes(word)) return word.toUpperCase()
  if (word === 'deepseek') return 'DeepSeek'
  return word.charAt(0).toUpperCase() + word.slice(1)
}
export function devinFamilyLabel(family: string): string {
  return family === 'swe-2' ? 'SWE-2' : family === 'swe-2-lightning' ? 'SWE-2 (Lightning)' : 'SWE-2 (Priority)'
}
export function devinChoices(ids: string[], current: string, effort = '') {
  const seen = new Set<string>()
  return ids.flatMap(id => {
    const family = devinFamily(id)
    if (!family) return [{ value: id, label: devinLabel(id, id === current ? effort : ''), desc: id }]
    if (seen.has(family[1])) return []
    seen.add(family[1])
    const members = ids.filter(candidate => devinFamily(candidate)?.[1] === family[1])
    const value = members.includes(current) ? current : id
    return [{ value, label: devinLabel(value, value === current ? effort : ''), desc: members.join(' · ') }]
  })
}
