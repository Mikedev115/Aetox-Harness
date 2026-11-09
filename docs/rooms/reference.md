# Room pack reference

## Where a pack lives

`<DataRoot>/rooms/<id>/`, one folder per room, named by the room's id.

The ask that starts a room names the folder the running app reads. Use that path
and nothing else: an installed app, a `wails dev` build and an `AETOX_DATA_ROOT`
override can each have a different data root, so a path remembered from another
build is a room nobody sees. If the ask does not name one, ask the person for
the folder shown in Settings — do not guess from the build you think you are in.

## room.json

| Field | Required | Rule |
|:--|:--|:--|
| `api` | yes | `1`. Any other value is refused (`room-api-unknown`). |
| `contract` | yes | `1`. Same rule, so a pack written for another version fails visibly. |
| `kind` | yes | `room`. |
| `id` | yes | Equal to the folder name, `^[a-z0-9][a-z0-9-]{1,31}$`. A folder starting `builtin-` is reserved (`room-id-reserved`). |
| `name` | yes | Shown on the card. Empty is refused (`room-name-missing`). |
| `description` | no | The second line of the card row. |
| `tier` | yes | `data` or `code`. Anything else is refused (`room-tier-unknown`). |
| `entry` | code only | A plain file name at the root of the pack, nothing else. Missing is refused (`room-entry-missing`). A `data` pack has no page, so this field is ignored for it. |
| `levels` | no | `company`, `dept`, `team`. Absent means all three. Any other word refuses the pack (`room-levels-unknown`). |
| `avatars` | no | `false` sends no roster at all, so the room draws no people. Absent means the app sends them. |
| `network` | no | Read and reported on the card. Nothing opens the network yet. |

A folder with no readable manifest is listed with `room-manifest-bad`, and a `kind`
that is not `room` with `room-kind-unknown`.

## Data packs: look.json

```json
{ "contract": 1, "palette": { "ground": "#dfe7ea", "wallIn": "#f8fbfc" } }
```

- `palette` keys are the names below and nothing else. A key the office does not use
  is dropped before it reaches the renderer.
- Values are `#rgb`, `#rrggbb` or `#rrggbbaa`. Anything else is dropped.
- The palette applies only at the rungs the pack declares in `levels`, and only while
  that room is chosen. A pack with no readable `look.json` draws in the office's own
  colours.
- `furniture` in this file is not read: a `data` pack changes colour, not the room.

The 41 names the office paints with:

| Surface | Names |
|:--|:--|
| ground | `ground`, `groundNight` |
| slabs and floors | `slab`, `slabDark`, `wood`, `woodDark`, `oak`, `walnut`, `terrazzo`, `stone`, `carpet`, `hub`, `bridge` |
| walls | `wallIn`, `wallOut`, `cap`, `glass`, `frame` |
| neutrals | `white`, `offwhite`, `grey`, `greyDark`, `charcoal`, `ink`, `metal`, `steel` |
| plants | `leaf`, `leafDark`, `leafLight`, `trunk`, `pot`, `terracotta` |
| signals | `cyan`, `blue`, `blueDeep`, `amber`, `green`, `red` |
| seating | `sofa`, `sofaLight`, `cushion` |

## Code packs: the page

- Served at `/aetox-room/<id>/<any file in the pack>`. The folder is the whole
  address space: a file beside `index.html` is reached by its name, nothing else is.
- The frame is `sandbox="allow-scripts"`. Scripts run, and the document has no origin
  of its own: `localStorage`, `sessionStorage`, `indexedDB` and cookies are refused.
  `fetch` and `XMLHttpRequest` reach the pack's own folder and nothing else: no app,
  no other pack, no internet.
- The document's policy is `default-src 'none'; script-src 'self' 'unsafe-inline';
  style-src 'self' 'unsafe-inline'; img-src 'self' data: blob:; font-src 'self';
  connect-src <the pack's own address>; worker-src blob:`. One file with inline script
  and inline style is enough. Images, models and data come from the pack folder, a
  data URL or a blob; a font comes from the pack folder; nothing else is fetched.
- A file is typed for the page: `.glb` is `model/gltf-binary`, `.gltf` is
  `model/gltf+json`, `.ktx2` is `image/ktx2`, and images, fonts, scripts, JSON and
  audio all arrive with their own type. A model beside the page loads with the
  renderer the pack carries.
- The frame covers the whole office panel: a tall column beside the chat in the
  desktop window, and the person can change its size. The panel's own background is
  `#0b0a12` until the page paints.

## The handshake

The page asks, and the app answers to the frame that asked.

| Direction | Message |
|:--|:--|
| pack to app | `{ type: 'ready' }` on load, and again when the pack wants the set again. It may carry `api: 1`; another value is refused and the card says the pack was made for a different Aetox. |
| app to pack | `{ type: 'hello', level, size: { w, h }, words }` when the frame loads |
| app to pack | `{ type: 'hello', api, level, words: { ..., andMore } }` in answer to `ready` |
| app to pack | `{ type: 'roster', level, agents: [...] }` |
| app to pack | `{ type: 'state', agents: [{ id, state }] }` about every 1.5 seconds |
| pack to app | `{ type: 'seen', n }` after taking a roster |

Four things about this traffic that are easy to get wrong:

- `hello` arrives twice, once with `size` and once with `api` and `andMore`. Merge
  the two messages. `size` is not repeated on the second one.
- `size` is the slot in CSS pixels at the moment of load. There is no second `hello`
  when the person resizes the panel, so the page listens for `resize` itself.
- `roster` can arrive twice within a moment (in answer to `ready`, then again on
  load). One of the two carries `level` and the other may not. Place people by `id`
  rather than appending them, and the second one is harmless.
- `state` repeats every ~1.5 seconds with every agent, not only the ones that
  changed. The app sends a line on the card when no `seen` arrives within 4 seconds
  of the frame loading, which is how a person learns the pack did not take its
  people. Answer with the number drawn, not with a promise.

Opened from the file system or in a plain browser tab there is no host: no `hello`,
no `roster`, no `state`. A room that draws people with nothing arriving is drawing
its own cast, which is the one thing this contract does not allow. Test a pack in the
app, at the rung it declares.

One agent:

| Field | What it is |
|:--|:--|
| `id` | Stable for that person. Use it as the key. |
| `name` | What the room writes under the figure, from the app. |
| `rank` | The app's own word for the rung. Draw it small or not at all. |
| `state` | One of the states below. |
| `sprite` | A PNG data URL of one baked frame: the app's own picture of that person, standing idle with its eyes open. It is re-sent only when the set of people changes, so cache it by `id`. It can be absent for a person the app has no picture of yet. The room draws this frame, and no character of its own. |
| `pose` | What the person is doing or feeling now, as the app's own key: `coding-t120`, `presenting-t90`, `asking`, or a feeling like `mood-happy` (the app composes the two in one key, so the first word says which it is). Absent while a figure is walking. |
| `poseSprite` | The frame `pose` names, a PNG data URL in the same shape as `sprite`. It arrives in `roster`, and again on a `state` tick when the key changed, once and not every tick. Cache it by key; the same key is the same picture. |
| `h` | How tall the figure is inside that image: 0 at the top edge, 1 at the bottom. The frame is square and keeps a margin around the rig, so `h` is under 1, about 0.8. |
| `feet` | Where the soles sit in that image: 0 at the top, 1 at the bottom. Scale the image so this line lands on the floor. |
| `walk` | The app's own walk cycle: one PNG strip of its eight baked frames, for the heading that faces the viewer. It arrives in `roster`, or on a later `state` tick once the cycle is baked; absent means it is not baked yet, so draw the step yourself until it comes. |
| `walkFeet` | Where the soles sit in that walk strip: the same floor line as `feet`, measured off a walk frame. |

The states the app reports:

| State | What it means in the room |
|:--|:--|
| `working` | at a desk, hands moving |
| `researching` | reading, screen light |
| `blocked` | stopped, asking, waiting on the person |
| `talking` | turned to somebody |
| `handoff` | passing a job on |
| `meeting` | gathered with others |
| `idle` | breathing, nothing assigned |
| `done` | finished, pleased |

The window may also send `moving` or `traveling` while it narrates a walk. `pose` is
what the person is doing now and `poseSprite` is that picture: an act (`typing`,
`search`, `presenting`) or a feeling (`mood-happy`, `mood-worried`), and the key's
first word says which. What never arrives is the cycle: a walk's other frames, a
pose's other phases, a blink. The pack draws the motion (the bob, the turn, the step)
from the one still it has, and draws no second character.

`words`:

| Key | Use |
|:--|:--|
| `hub`, `lounge`, `core`, `lobby`, `meeting`, `knowledge`, `delivery`, `memory` | the room's areas, as signs |
| `team`, `company`, `tagline`, `upperFloor` | the rung and the labels the app writes on its own floors |
| `jobs` | the word after a count of finished jobs |
| `andMore` | "and N more", for a list the room shortens |

Every label on the page comes from here. A pack that writes Thai or English of its
own is wrong in one of the two windows.

## Standing a picture on the floor

`sprite` is square, and the rig is drawn inside it with a margin, so the figure fills
about four fifths of the image (`h`). To stand a person `tall` pixels high:

```
H = tall / h            // draw the image H × H
top = floorY - feet * H // its top edge, so the soles land on the floor line
```

Scaling the image by `feet` alone draws the figure `h` times as tall as intended,
about a fifth short: small people around normal furniture. The height itself is the
pack's to choose, from something every room has: a desk comes to a standing person's
waist, a doorway clears their head. Pick that one length and stand every figure
against it, so the room has a scale rather than a pile of separate guesses.

## Walking

A person crossing the room is the pack's drawing. `walk` is the app's own cycle: a
strip of eight frames, the frame chosen by the ground covered and the direction by
whether the strip is mirrored. A room that has no strip yet — it is baked after the
standing frame, so the first seconds of a turn can go without one — draws the step
itself, and three lines are what separate a walk from a glitch.

- Mirror, do not slide. Moving against the way the picture faces means turning the
  picture, not the path: flip the frame on the way left, and hold its size for the
  whole walk, because a figure that grows on the way to its seat reads as a
jump.
- Steps, not hops. Two small rises a second, at most a fortieth of the figure's
  height, timed by the ground covered rather than by the clock. A tenth of the height
  twice a second is a pogo stick.
- The soles own the floor. `feet` puts them on the floor line, the shadow is the
  contact, and every frame of the walk keeps that one line: the picture moves, the
  place the person stands on does not.

A `state` of `moving` or `traveling` is the window narrating a walk, and `pose` is
absent through it: the walk strip is what a walk is drawn from, and the standing
frame is what is drawn when no strip has arrived.

## What the app draws over the page

- Top left: where the room stands, the department and then the team, as a button.
- Top right: the room picker and the count of events.
- The head, the card, the hint bar, the event log and the empty message sit above the
  page, whatever the pack paints under them.

Those two top corners belong to the app. The rest of the frame is the room's.
