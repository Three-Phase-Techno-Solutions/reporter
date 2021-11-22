package report

const (
	/*
		"dbDSN TODO fill this in directly or through environment variable
		Build a DSN e.g. postgres://username:password@url.com:5432/dbName"
	*/

	//dbDSN = `postgres://db_reader:5TVezcF%YZ@tpts.3phtechsolutions.com/eyJrIjoiNHM2d0taYlZ5b1pkeWxFSm9yMk5SSGZwNmRrd3kwTzUiLCJuIjoiR3JhZmFuYSBSZXBvcnRlciIsImlkIjoyfQ==/thingsboard`
	//dbDSN = `postgres://db_reader:0110401611@localhost:5432/thingsboard`

	strTemplateDefault = `
		\documentclass{article}
		\renewcommand{\rmdefault}{cmss}
		\usepackage[utf8]{inputenc}
		\DeclareUnicodeCharacter{2218}{\circ}
		\usepackage[usegeometry]{typearea}
		\usepackage[a4paper]{geometry}
		\geometry{
			margin = 1cm,
			headheight =2cm,
			includehead,
			includefoot,
		}
		\usepackage{float}
		\usepackage{longtable}
		\setlength\tabcolsep{3pt}
		\usepackage{pbox}
		\usepackage{graphicx}
		\usepackage{subcaption}
		\usepackage{fancyhdr}
		\usepackage{lastpage}
		\usepackage[dvipsnames]{xcolor}
		\usepackage[ddmmyyyy]{datetime}

		\settimeformat{hhmmsstime}
		\renewcommand{\dateseparator}{-}
		\pagestyle{fancyplain}

		\title{[[.TitleReport]] [[if .VariableValues]] \\ \large [[.VariableValues]] [[end]] [[if .Description]] \\ \small [[.Description]] [[end]]}
		\date{[[.TimeFrom]]\\to\\[[.TimeTo]]}
		\graphicspath{{images/}}

		\begin{document}
			\centering
			\fancyhf{}
			\rfoot{Page \thepage\hspace{1pt} of \pageref{LastPage}}
			\lfoot{Report generated on \today, at \currenttime}
			\lhead{\includegraphics[width=4cm]{[[.NameLogoCompany]].png}}
			\chead{
				\LARGE{\textbf{[[.TitleReport]]}}\\ 
				\Large{\textbf{[[.TitleDashBoard]]}}\\ 
				\scriptsize{Time Period: [[.TimeFrom]] to [[.TimeTo]]}
			}
			\rhead{\includegraphics[width=2cm]{[[.NameLogoClient]].png}}
			%\renewcommand{\headrulewidth}{0.4pt}% Default \headrulewidth is 0.4pt
			\renewcommand{\footrulewidth}{0.4pt}% Default \footrulewidth is 0pt

			\storeareas\MySavedValues
			[[$widthLinePortraitMax := 70]]
			[[$widthLineLandscapeMax := 132]]
			[[range .Panels]]
				\begin{figure}[H]
				\vspace{5mm}
					\includegraphics[width=\textwidth]{[[.NameImage]].png}
					\caption{[[.TitlePanel]]}
				\end{figure}
				[[$widthTextPerLine := .DataTableMC.WidthTextPerLine]]
				[[if gt .DataTableMC.CountCols 0]]
					[[$countColSeparator := .DataTableMC.CountColsSeperator]]
					[[if gt $widthTextPerLine $widthLinePortraitMax]]
						\clearpage
						\KOMAoptions{paper=landscape,DIV=20}
						\newgeometry{
							margin = 1cm,
							headheight =2cm,
							includehead,
							includefoot,
						}
						\fontsize{10}{12}\selectfont
						[[if gt $widthTextPerLine $widthLineLandscapeMax]]
							\fontsize{[[.DataTableMC.FontSizeTableLandScape]]}{[[.DataTableMC.LineSpacingLandscape]]}\selectfont
						[[end]]
						\fancyheadoffset{0pt}% recalculate headwidth for fancy

					[[end]]
					\begin{longtable}[]{[[range .DataTableMC.RowHeader]]|c[[end]]|}
						\hline
						[[range $i, $v := .DataTableMC.RowHeader]]
							[[if gt $widthTextPerLine $widthLinePortraitMax]]
								\pbox[c][][c]{30pt}{\centering{\bf{[[$v.Value]]}}}
							[[else]]
								\bf [[$v.Value]]
							[[end]]
							[[if gt $countColSeparator $i]] & [[end]]
						[[end]]\\
						\endfirsthead
						\caption{Data Table - [[.TitlePanel]] \small{(...continued)}}\\
						\hline
						[[range $i, $v := .DataTableMC.RowHeader]]
							[[if gt $widthTextPerLine $widthLinePortraitMax]]
								\pbox[c][][c]{30pt}{\centering{\bf{[[$v.Value]]}}}
							[[else]]
								\bf [[$v.Value]]
							[[end]]
							[[if gt $countColSeparator $i]] & [[end]]
						[[end]]\\ 
						\endhead
						\multicolumn{[[.DataTableMC.CountCols]]}{r}{\small{(continued...)}} \\
						\endfoot
						\endlastfoot
										
						[[range .DataTableMC.Rows]]
						
							\hline
						
							[[range $i, $v := .Row]] \textcolor[HTML]{[[$v.ColorText]]}{[[$v.Value]]} [[if gt $countColSeparator $i]] & [[end]] [[end]]\\

						[[end]]
						\hline
						\caption{Data Table - [[.TitlePanel]]}\\
					\end{longtable}
					\center{-}
				[[end]]
				[[if gt $widthTextPerLine $widthLinePortraitMax]]
					\clearpage
					\fontsize{10}{12}\selectfont
					\MySavedValues
				[[end]]
			[[end]]
		\end{document}
	`

	strTemplateGrid = `
		%use square brackets as golang text templating delimiters
		\documentclass{article}
		\usepackage{graphicx}
		\usepackage[margin=0.5in]{geometry}

		\graphicspath{ {images/} }
		\begin{document}
		\title{[[.TitleReport]] [[if .VariableValues]] \\ \large [[.VariableValues]] [[end]] [[if .Description]] \\ \small [[.Description]] [[end]]}
		\date{[[.TimeFrom]]\\to\\[[.TimeTo]]}
		\maketitle
		\begin{center}
		[[range .Panels]][[if .IsPartialWidth]]\begin{minipage}{[[.Width]]\textwidth}
		\includegraphics[width=\textwidth]{image[[.ID]]}
		\end{minipage}
		[[else]]\par
		\vspace{0.5cm}
		\includegraphics[width=\textwidth]{image[[.ID]]}
		\par
		\vspace{0.5cm}
		[[end]][[end]]
		\end{center}
		\end{document}
`
	strTemplateGridTPTS = `
		\documentclass{article}
		\renewcommand{\rmdefault}{cmss}
		\usepackage[utf8]{inputenc}
		\DeclareUnicodeCharacter{2218}{\circ}
		\usepackage[usegeometry]{typearea}
		\usepackage{geometry}
		\geometry{
			a4paper,
			margin = 1cm,
			headheight =2cm,
			includehead,
			includefoot,
		}
		\usepackage{longtable}
		\usepackage{graphicx}
		\usepackage{subcaption}
		\usepackage{fancyhdr}
		\usepackage{lastpage}
		\usepackage[dvipsnames]{xcolor}
		\usepackage[ddmmyyyy]{datetime}

		\settimeformat{hhmmsstime}
		\renewcommand{\dateseparator}{-}
		\pagestyle{fancyplain}

		\title{[[.TitleReport]]}
		\date{\today}
		\graphicspath{{images/}}

		\begin{document}
			\centering
			\fancyhf{}
			\rfoot{Page \thepage\hspace{1pt} of \pageref{LastPage}}
			\lfoot{Report generated on \today, at \currenttime}
			\lhead{\includegraphics[width=4cm]{[[.NameLogoCompany]].png}}
			\chead{
				\LARGE{\textbf{[[.TitleReport]]}}\\ 
				\Large{\textbf{[[.TitleDashBoard]]}}\\ 
				\scriptsize{Time Period: [[.TimeFrom]] to [[.TimeTo]]}
			}
			\rhead{\includegraphics[width=2cm]{[[.NameLogoClient]].png}}
			%\renewcommand{\headrulewidth}{0.4pt}% Default \headrulewidth is 0.4pt
			\renewcommand{\footrulewidth}{0.4pt}% Default \footrulewidth is 0pt
			\KOMAoptions{paper=landscape,DIV=20}
			\newgeometry{
				margin = 1cm,
				headheight =2cm,
				includehead,
				includefoot,
			}
			\fancyheadoffset{0pt}% recalculate headwidth for fancy
			\begin{center}
				[[range .Panels]][[if .IsPartialWidth]]\begin{minipage}{[[.Width]]\textwidth}
				\includegraphics[width=\textwidth]{image[[.ID]]}
				\end{minipage}
				[[else]]\par
				\vspace{0.5cm}
				\includegraphics[width=\textwidth]{image[[.ID]]}
				\par
				\vspace{0.5cm}
				[[end]][[end]]

			\end{center}
		\end{document}
`
)

var mapUnits = map[string]string{
	"Cost":          "Rs.",
	"Amount":        "Rs.",
	"Energy Cost":   "Rs.",
	"Power":         "kW",
	"Voltage":       "V",
	"Phase Voltage": "V",
	"Line Voltage":  "V",
	"Current":       "A",
	"Energy":        "kWh",
	"Percentage":    `\%`,
	"Temperature":   `$^{\circ}$C`,
	"Water Usage":   "m$^{3}$",
	"Gas Usage":     "m$^{3}$",
	"Frequency":     "Hz",
	"kVArh":         "kVArh",
	"kVAh":          "kVAh",
	"kWh":           "kWh",
	"kVAr":          "kVAr",
	"kVA":           "kVA",
	"kW":            "kW",
	"co2":           "t",
	"IR":            "A",
	"IY":            "A",
	"IB":            "A",
	"RY":            "V",
	"RB":            "V",
	"YB":            "V",
	"VRN":           "V",
	"VYN":           "V",
	"VBN":           "V",
}

var panelsToIgnore = map[string]bool{
	"Featured":           false,
	"Device":             false,
	"Beer Glass Analogy": false,
	"Menu":               false,
	"CO2":                false,
}
