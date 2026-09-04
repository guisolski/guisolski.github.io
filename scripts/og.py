#!/usr/bin/env python3
"""Generate assets/images/og.png, the 1200x630 share card.

The card is the page's own opening in miniature: the dawn spectrum as a rule,
the name, what he does, and where he is. Regenerate it with `make og` after
changing any of those strings — it is a build artifact, not hand-drawn art.
"""

from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

W, H = 1200, 630
PAPER = (255, 255, 255)
INK = (29, 30, 66)
MUTED = (97, 99, 128)
ACCENT = (143, 87, 0)

# The dawn walk: night indigo -> dusk rose -> star amber, the same spectrum
# the career rail travels on the page.
DAWN = [(75, 60, 224), (224, 75, 166), (245, 166, 35)]

ROOT = Path(__file__).resolve().parents[1]
SANS = "/System/Library/Fonts/HelveticaNeue.ttc"
MONO = "/System/Library/Fonts/Menlo.ttc"

NAME = "Guilherme Solski Alves"
EYEBROW = "BACKEND ENGINEER · GO · REMOTE, UTC−3"
LEAD = "Backend engineer with 7+ years on high-throughput"
LEAD2 = "production systems, and co-founder of FINK AI."
FOOT = "guisolski.github.io · Curitiba, Paraná, Brazil"


def lerp(a, b, t):
    return tuple(round(x + (y - x) * t) for x, y in zip(a, b))


def dawn_at(t):
    """Sample the three-stop spectrum at 0..1."""
    if t <= 0.5:
        return lerp(DAWN[0], DAWN[1], t / 0.5)
    return lerp(DAWN[1], DAWN[2], (t - 0.5) / 0.5)


def font(path, size, index=0):
    return ImageFont.truetype(path, size, index=index)


def aurora(img):
    """The page's aurora, flattened: three soft washes off the text column."""
    glow = Image.new("RGB", (W, H), PAPER)
    px = glow.load()
    spots = [
        ((0.16, 0.18), 0.42, DAWN[0], 0.16),
        ((0.86, 0.10), 0.34, DAWN[2], 0.20),
        ((0.62, 0.86), 0.30, DAWN[1], 0.12),
    ]
    for y in range(H):
        for x in range(W):
            r, g, b = PAPER
            for (cx, cy), radius, colour, strength in spots:
                dx = (x / W - cx) * 1.0
                dy = (y / H - cy) * (H / W)
                d = (dx * dx + dy * dy) ** 0.5
                if d >= radius:
                    continue
                falloff = (1 - d / radius) ** 2 * strength
                r += (colour[0] - r) * falloff
                g += (colour[1] - g) * falloff
                b += (colour[2] - b) * falloff
            px[x, y] = (round(r), round(g), round(b))
    img.paste(glow, (0, 0))


def dawn_rule(draw, x0, y0, x1, height):
    for i in range(x1 - x0):
        draw.rectangle(
            [x0 + i, y0, x0 + i + 1, y0 + height],
            fill=dawn_at(i / max(1, x1 - x0 - 1)),
        )


def main():
    img = Image.new("RGB", (W, H), PAPER)
    aurora(img)
    draw = ImageDraw.Draw(img)

    pad = 88
    dawn_rule(draw, pad, 110, W - pad, 4)

    # Face 1 of HelveticaNeue.ttc is Bold; 0 is Regular. Pillow will happily
    # take an index that is Italic instead, so the numbers are pinned here.
    draw.text((pad, 170), EYEBROW, font=font(MONO, 21), fill=ACCENT)
    draw.text((pad, 228), NAME, font=font(SANS, 76, index=1), fill=INK)
    draw.text((pad, 348), LEAD, font=font(SANS, 34), fill=INK)
    draw.text((pad, 396), LEAD2, font=font(SANS, 34), fill=INK)
    draw.text((pad, 492), FOOT, font=font(MONO, 22), fill=MUTED)

    out = ROOT / "assets" / "images" / "og.png"
    out.parent.mkdir(parents=True, exist_ok=True)
    img.save(out, optimize=True)
    print(f"wrote {out.relative_to(ROOT)} ({out.stat().st_size // 1024} KiB)")


if __name__ == "__main__":
    main()
