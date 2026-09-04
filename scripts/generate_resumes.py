#!/usr/bin/env python3
"""Generate ATS-friendly EN/PT résumés for guisolski.github.io.

Uses only verified career facts from the personal site / CVs.
"""

from pathlib import Path

from reportlab.lib.pagesizes import letter
from reportlab.lib.styles import ParagraphStyle, getSampleStyleSheet
from reportlab.lib.units import inch
from reportlab.platypus import Paragraph, SimpleDocTemplate, Spacer

ROOT = Path(__file__).resolve().parents[1]
EN_OUT = ROOT / "assets/pdf/resume.pdf"
PT_OUT = ROOT / "assets/pdf/Curriculum/portugues.pdf"


def styles():
    base = getSampleStyleSheet()
    return {
        "name": ParagraphStyle(
            "Name",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=16,
            leading=20,
            spaceAfter=4,
        ),
        "contact": ParagraphStyle(
            "Contact",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=9,
            leading=12,
            spaceAfter=10,
        ),
        "section": ParagraphStyle(
            "Section",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=11,
            leading=14,
            spaceBefore=10,
            spaceAfter=4,
        ),
        "body": ParagraphStyle(
            "Body",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=9.5,
            leading=12.5,
            spaceAfter=4,
        ),
        "job": ParagraphStyle(
            "Job",
            parent=base["Normal"],
            fontName="Helvetica-Bold",
            fontSize=10,
            leading=13,
            spaceBefore=6,
            spaceAfter=1,
        ),
        "meta": ParagraphStyle(
            "Meta",
            parent=base["Normal"],
            fontName="Helvetica-Oblique",
            fontSize=9,
            leading=11,
            spaceAfter=2,
        ),
        "bullet": ParagraphStyle(
            "Bullet",
            parent=base["Normal"],
            fontName="Helvetica",
            fontSize=9.5,
            leading=12.5,
            leftIndent=12,
            spaceAfter=2,
        ),
    }


def bullet(text, s):
    return Paragraph(f"• {text}", s["bullet"])


def build_en(path: Path):
    s = styles()
    doc = SimpleDocTemplate(
        str(path),
        pagesize=letter,
        leftMargin=0.7 * inch,
        rightMargin=0.7 * inch,
        topMargin=0.55 * inch,
        bottomMargin=0.55 * inch,
        title="Guilherme Solski Alves — Résumé",
        author="Guilherme Solski Alves",
    )
    story = [
        Paragraph("Guilherme Solski Alves", s["name"]),
        Paragraph(
            "Curitiba, Paraná, Brazil — Open to remote and relocation<br/>"
            "+55 41 99628-6624 · guilhermesolskialves@gmail.com<br/>"
            "LinkedIn: linkedin.com/in/guilherme-solski-alves · "
            "GitHub: github.com/guisolski · Site: guisolski.github.io",
            s["contact"],
        ),
        Paragraph("Summary", s["section"]),
        Paragraph(
            "Backend engineer with 7+ years in production systems. Currently writing Go at "
            "Mercado Libre for e-commerce platforms serving 1M+ customers, with end-to-end "
            "ownership of observability (OpenTelemetry; New Relic, Grafana, Datadog, Kibana). "
            "Part-time co-founder of FINK AI. Open to full-time backend roles (remote or relocation).",
            s["body"],
        ),
        Paragraph("Experience", s["section"]),
        Paragraph("Mercado Libre — Software Engineer (Go)", s["job"]),
        Paragraph("Oct 2022 – Present · E-commerce platforms", s["meta"]),
        bullet(
            "Write Go for scalable services serving 1M+ customers; use goroutines and "
            "in-memory caching where appropriate.",
            s,
        ),
        bullet(
            "Own the observability chain end to end: application metrics, OpenTelemetry "
            "instrumentation, and operational dashboards and alerts across New Relic, Grafana, "
            "Datadog, and Kibana.",
            s,
        ),
        Paragraph("FINK AI — Co-Founder (part-time)", s["job"]),
        Paragraph("May 2025 – Present · Alongside full-time employment", s["meta"]),
        bullet(
            "Personal-finance product with 20,000+ registered users and 4.2M+ financial records; "
            "Open Finance.",
            s,
        ),
        Paragraph("ExxonMobil — Java Developer", s["job"]),
        Paragraph("Jul 2021 – Oct 2022", s["meta"]),
        bullet("Azure with Terraform; Confluent Cloud Kafka + Avro.", s),
        bullet(
            "MSP transition program; Leadership Recognition Award (2022).",
            s,
        ),
        Paragraph("Composable Bit — Java Developer", s["job"]),
        Paragraph("May 2021 – Jul 2021 · Self-employed", s["meta"]),
        Paragraph("Go44 — Ruby Developer (part-time)", s["job"]),
        Paragraph("Oct 2020 – 2021 · Industry 4.0 systems", s["meta"]),
        Paragraph("SaaSTec Labs — PHP Developer / Intern", s["job"]),
        Paragraph("Mar 2020 – 2021 · E-commerce systems", s["meta"]),
        Paragraph("Banco do Brasil — Intern (Java)", s["job"]),
        Paragraph("Jan 2019 – Mar 2020 · Web systems (Java, JavaScript, MySQL)", s["meta"]),
        Paragraph("Education", s["section"]),
        bullet("MBA, Data Science & Analytics — USP / ESALQ, 2022–2024", s),
        bullet("Bachelor of Computer Science — PUCPR, 2018–2021", s),
        bullet("Digital Game Developer (technical) — IFPR, 2015–2017", s),
        Paragraph("Languages", s["section"]),
        Paragraph(
            "Portuguese — Native · English — Full professional · Spanish — Professional working",
            s["body"],
        ),
        Paragraph("Technical skills", s["section"]),
        Paragraph(
            "Go, Java, Python, SQL, JavaScript · OpenTelemetry · New Relic, Grafana, Datadog, "
            "Kibana · Azure, Terraform · Kafka / Avro · Caching, concurrent services",
            s["body"],
        ),
    ]
    doc.build(story)


def build_pt(path: Path):
    s = styles()
    doc = SimpleDocTemplate(
        str(path),
        pagesize=letter,
        leftMargin=0.7 * inch,
        rightMargin=0.7 * inch,
        topMargin=0.55 * inch,
        bottomMargin=0.55 * inch,
        title="Guilherme Solski Alves — Currículo",
        author="Guilherme Solski Alves",
    )
    story = [
        Paragraph("Guilherme Solski Alves", s["name"]),
        Paragraph(
            "Curitiba, Paraná, Brasil — Aberto a remoto e realocação<br/>"
            "+55 41 99628-6624 · guilhermesolskialves@gmail.com<br/>"
            "LinkedIn: linkedin.com/in/guilherme-solski-alves · "
            "GitHub: github.com/guisolski · Site: guisolski.github.io",
            s["contact"],
        ),
        Paragraph("Resumo", s["section"]),
        Paragraph(
            "Engenheiro de backend com mais de 7 anos em sistemas de produção. Atualmente "
            "escreve Go no Mercado Livre para plataformas de e-commerce com 1M+ clientes, "
            "com ownership ponta a ponta de observabilidade (OpenTelemetry; New Relic, Grafana, "
            "Datadog, Kibana). Co-fundador em meio período da FINK AI. Aberto a vagas "
            "full-time de backend (remoto ou realocação).",
            s["body"],
        ),
        Paragraph("Experiência", s["section"]),
        Paragraph("Mercado Livre — Software Engineer (Go)", s["job"]),
        Paragraph("Out 2022 – Atual · Plataformas de e-commerce", s["meta"]),
        bullet(
            "Desenvolvimento em Go de serviços escaláveis para 1M+ clientes; uso de "
            "goroutines e cache em memória quando apropriado.",
            s,
        ),
        bullet(
            "Ownership da cadeia de observabilidade: métricas de aplicação, instrumentação "
            "OpenTelemetry e dashboards/alertas operacionais em New Relic, Grafana, Datadog e Kibana.",
            s,
        ),
        Paragraph("FINK AI — Co-fundador (meio período)", s["job"]),
        Paragraph("Mai 2025 – Atual · Em paralelo ao emprego full-time", s["meta"]),
        bullet(
            "Produto de finanças pessoais com 20.000+ usuários registrados e 4,2M+ registros "
            "financeiros; Open Finance.",
            s,
        ),
        Paragraph("ExxonMobil — Desenvolvedor Java", s["job"]),
        Paragraph("Jul 2021 – Out 2022", s["meta"]),
        bullet("Azure com Terraform; Confluent Cloud Kafka + Avro.", s),
        bullet(
            "Programa de transição MSP; Leadership Recognition Award (2022).",
            s,
        ),
        Paragraph("Composable Bit — Desenvolvedor Java", s["job"]),
        Paragraph("Mai 2021 – Jul 2021 · Autônomo", s["meta"]),
        Paragraph("Go44 — Desenvolvedor Ruby (meio período)", s["job"]),
        Paragraph("Out 2020 – 2021 · Sistemas Industry 4.0", s["meta"]),
        Paragraph("SaaSTec Labs — Desenvolvedor PHP / Estagiário", s["job"]),
        Paragraph("Mar 2020 – 2021 · Sistemas de e-commerce", s["meta"]),
        Paragraph("Banco do Brasil — Estagiário (Java)", s["job"]),
        Paragraph("Jan 2019 – Mar 2020 · Sistemas web (Java, JavaScript, MySQL)", s["meta"]),
        Paragraph("Formação", s["section"]),
        bullet("MBA em Data Science & Analytics — USP / ESALQ, 2022–2024", s),
        bullet("Bacharelado em Ciência da Computação — PUCPR, 2018–2021", s),
        bullet("Técnico em Desenvolvimento de Jogos Digitais — IFPR, 2015–2017", s),
        Paragraph("Idiomas", s["section"]),
        Paragraph(
            "Português — Nativo · Inglês — Proficiência profissional completa · "
            "Espanhol — Proficiência profissional de trabalho",
            s["body"],
        ),
        Paragraph("Competências técnicas", s["section"]),
        Paragraph(
            "Go, Java, Python, SQL, JavaScript · OpenTelemetry · New Relic, Grafana, Datadog, "
            "Kibana · Azure, Terraform · Kafka / Avro · Cache, serviços concorrentes",
            s["body"],
        ),
    ]
    doc.build(story)


def main():
    EN_OUT.parent.mkdir(parents=True, exist_ok=True)
    PT_OUT.parent.mkdir(parents=True, exist_ok=True)
    build_en(EN_OUT)
    build_pt(PT_OUT)
    print(f"wrote {EN_OUT} ({EN_OUT.stat().st_size} bytes)")
    print(f"wrote {PT_OUT} ({PT_OUT.stat().st_size} bytes)")


if __name__ == "__main__":
    main()
