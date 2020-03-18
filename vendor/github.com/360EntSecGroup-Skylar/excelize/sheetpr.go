// Copyright 2016 - 2020 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX files. Support reads and writes XLSX file generated by
// Microsoft Excel™ 2007 and later. Support save file without losing original
// charts of XLSX. This library needs Go version 1.10 or later.

package excelize

// SheetPrOption is an option of a view of a worksheet. See SetSheetPrOptions().
type SheetPrOption interface {
	setSheetPrOption(view *xlsxSheetPr)
}

// SheetPrOptionPtr is a writable SheetPrOption. See GetSheetPrOptions().
type SheetPrOptionPtr interface {
	SheetPrOption
	getSheetPrOption(view *xlsxSheetPr)
}

type (
	// CodeName is a SheetPrOption
	CodeName string
	// EnableFormatConditionsCalculation is a SheetPrOption
	EnableFormatConditionsCalculation bool
	// Published is a SheetPrOption
	Published bool
	// FitToPage is a SheetPrOption
	FitToPage bool
	// AutoPageBreaks is a SheetPrOption
	AutoPageBreaks bool
	// OutlineSummaryBelow is an outlinePr, within SheetPr option
	OutlineSummaryBelow bool
)

// setSheetPrOption implements the SheetPrOption interface.
func (o OutlineSummaryBelow) setSheetPrOption(pr *xlsxSheetPr) {
	if pr.OutlinePr == nil {
		pr.OutlinePr = new(xlsxOutlinePr)
	}
	pr.OutlinePr.SummaryBelow = bool(o)
}

// getSheetPrOption implements the SheetPrOptionPtr interface.
func (o *OutlineSummaryBelow) getSheetPrOption(pr *xlsxSheetPr) {
	// Excel default: true
	if pr == nil || pr.OutlinePr == nil {
		*o = true
		return
	}
	*o = OutlineSummaryBelow(defaultTrue(&pr.OutlinePr.SummaryBelow))
}

// setSheetPrOption implements the SheetPrOption interface and specifies a
// stable name of the sheet.
func (o CodeName) setSheetPrOption(pr *xlsxSheetPr) {
	pr.CodeName = string(o)
}

// getSheetPrOption implements the SheetPrOptionPtr interface and get the
// stable name of the sheet.
func (o *CodeName) getSheetPrOption(pr *xlsxSheetPr) {
	if pr == nil {
		*o = ""
		return
	}
	*o = CodeName(pr.CodeName)
}

// setSheetPrOption implements the SheetPrOption interface and flag indicating
// whether the conditional formatting calculations shall be evaluated.
func (o EnableFormatConditionsCalculation) setSheetPrOption(pr *xlsxSheetPr) {
	pr.EnableFormatConditionsCalculation = boolPtr(bool(o))
}

// getSheetPrOption implements the SheetPrOptionPtr interface and get the
// settings of whether the conditional formatting calculations shall be
// evaluated.
func (o *EnableFormatConditionsCalculation) getSheetPrOption(pr *xlsxSheetPr) {
	if pr == nil {
		*o = true
		return
	}
	*o = EnableFormatConditionsCalculation(defaultTrue(pr.EnableFormatConditionsCalculation))
}

// setSheetPrOption implements the SheetPrOption interface and flag indicating
// whether the worksheet is published.
func (o Published) setSheetPrOption(pr *xlsxSheetPr) {
	pr.Published = boolPtr(bool(o))
}

// getSheetPrOption implements the SheetPrOptionPtr interface and get the
// settings of whether the worksheet is published.
func (o *Published) getSheetPrOption(pr *xlsxSheetPr) {
	if pr == nil {
		*o = true
		return
	}
	*o = Published(defaultTrue(pr.Published))
}

// setSheetPrOption implements the SheetPrOption interface.
func (o FitToPage) setSheetPrOption(pr *xlsxSheetPr) {
	if pr.PageSetUpPr == nil {
		if !o {
			return
		}
		pr.PageSetUpPr = new(xlsxPageSetUpPr)
	}
	pr.PageSetUpPr.FitToPage = bool(o)
}

// getSheetPrOption implements the SheetPrOptionPtr interface.
func (o *FitToPage) getSheetPrOption(pr *xlsxSheetPr) {
	// Excel default: false
	if pr == nil || pr.PageSetUpPr == nil {
		*o = false
		return
	}
	*o = FitToPage(pr.PageSetUpPr.FitToPage)
}

// setSheetPrOption implements the SheetPrOption interface.
func (o AutoPageBreaks) setSheetPrOption(pr *xlsxSheetPr) {
	if pr.PageSetUpPr == nil {
		if !o {
			return
		}
		pr.PageSetUpPr = new(xlsxPageSetUpPr)
	}
	pr.PageSetUpPr.AutoPageBreaks = bool(o)
}

// getSheetPrOption implements the SheetPrOptionPtr interface.
func (o *AutoPageBreaks) getSheetPrOption(pr *xlsxSheetPr) {
	// Excel default: false
	if pr == nil || pr.PageSetUpPr == nil {
		*o = false
		return
	}
	*o = AutoPageBreaks(pr.PageSetUpPr.AutoPageBreaks)
}

// SetSheetPrOptions provides a function to sets worksheet properties.
//
// Available options:
//   CodeName(string)
//   EnableFormatConditionsCalculation(bool)
//   Published(bool)
//   FitToPage(bool)
//   AutoPageBreaks(bool)
//   OutlineSummaryBelow(bool)
func (f *File) SetSheetPrOptions(name string, opts ...SheetPrOption) error {
	sheet, err := f.workSheetReader(name)
	if err != nil {
		return err
	}
	pr := sheet.SheetPr
	if pr == nil {
		pr = new(xlsxSheetPr)
		sheet.SheetPr = pr
	}

	for _, opt := range opts {
		opt.setSheetPrOption(pr)
	}
	return err
}

// GetSheetPrOptions provides a function to gets worksheet properties.
//
// Available options:
//   CodeName(string)
//   EnableFormatConditionsCalculation(bool)
//   Published(bool)
//   FitToPage(bool)
//   AutoPageBreaks(bool)
//   OutlineSummaryBelow(bool)
func (f *File) GetSheetPrOptions(name string, opts ...SheetPrOptionPtr) error {
	sheet, err := f.workSheetReader(name)
	if err != nil {
		return err
	}
	pr := sheet.SheetPr

	for _, opt := range opts {
		opt.getSheetPrOption(pr)
	}
	return err
}

type (
	// PageMarginBottom specifies the bottom margin for the page.
	PageMarginBottom float64
	// PageMarginFooter specifies the footer margin for the page.
	PageMarginFooter float64
	// PageMarginHeader specifies the header margin for the page.
	PageMarginHeader float64
	// PageMarginLeft specifies the left margin for the page.
	PageMarginLeft float64
	// PageMarginRight specifies the right margin for the page.
	PageMarginRight float64
	// PageMarginTop specifies the top margin for the page.
	PageMarginTop float64
)

// setPageMargins provides a method to set the bottom margin for the worksheet.
func (p PageMarginBottom) setPageMargins(pm *xlsxPageMargins) {
	pm.Bottom = float64(p)
}

// setPageMargins provides a method to get the bottom margin for the worksheet.
func (p *PageMarginBottom) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.75
	if pm == nil || pm.Bottom == 0 {
		*p = 0.75
		return
	}
	*p = PageMarginBottom(pm.Bottom)
}

// setPageMargins provides a method to set the footer margin for the worksheet.
func (p PageMarginFooter) setPageMargins(pm *xlsxPageMargins) {
	pm.Footer = float64(p)
}

// setPageMargins provides a method to get the footer margin for the worksheet.
func (p *PageMarginFooter) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.3
	if pm == nil || pm.Footer == 0 {
		*p = 0.3
		return
	}
	*p = PageMarginFooter(pm.Footer)
}

// setPageMargins provides a method to set the header margin for the worksheet.
func (p PageMarginHeader) setPageMargins(pm *xlsxPageMargins) {
	pm.Header = float64(p)
}

// setPageMargins provides a method to get the header margin for the worksheet.
func (p *PageMarginHeader) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.3
	if pm == nil || pm.Header == 0 {
		*p = 0.3
		return
	}
	*p = PageMarginHeader(pm.Header)
}

// setPageMargins provides a method to set the left margin for the worksheet.
func (p PageMarginLeft) setPageMargins(pm *xlsxPageMargins) {
	pm.Left = float64(p)
}

// setPageMargins provides a method to get the left margin for the worksheet.
func (p *PageMarginLeft) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.7
	if pm == nil || pm.Left == 0 {
		*p = 0.7
		return
	}
	*p = PageMarginLeft(pm.Left)
}

// setPageMargins provides a method to set the right margin for the worksheet.
func (p PageMarginRight) setPageMargins(pm *xlsxPageMargins) {
	pm.Right = float64(p)
}

// setPageMargins provides a method to get the right margin for the worksheet.
func (p *PageMarginRight) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.7
	if pm == nil || pm.Right == 0 {
		*p = 0.7
		return
	}
	*p = PageMarginRight(pm.Right)
}

// setPageMargins provides a method to set the top margin for the worksheet.
func (p PageMarginTop) setPageMargins(pm *xlsxPageMargins) {
	pm.Top = float64(p)
}

// setPageMargins provides a method to get the top margin for the worksheet.
func (p *PageMarginTop) getPageMargins(pm *xlsxPageMargins) {
	// Excel default: 0.75
	if pm == nil || pm.Top == 0 {
		*p = 0.75
		return
	}
	*p = PageMarginTop(pm.Top)
}

// PageMarginsOptions is an option of a page margin of a worksheet. See
// SetPageMargins().
type PageMarginsOptions interface {
	setPageMargins(layout *xlsxPageMargins)
}

// PageMarginsOptionsPtr is a writable PageMarginsOptions. See
// GetPageMargins().
type PageMarginsOptionsPtr interface {
	PageMarginsOptions
	getPageMargins(layout *xlsxPageMargins)
}

// SetPageMargins provides a function to set worksheet page margins.
//
// Available options:
//   PageMarginBotom(float64)
//   PageMarginFooter(float64)
//   PageMarginHeader(float64)
//   PageMarginLeft(float64)
//   PageMarginRight(float64)
//   PageMarginTop(float64)
func (f *File) SetPageMargins(sheet string, opts ...PageMarginsOptions) error {
	s, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	pm := s.PageMargins
	if pm == nil {
		pm = new(xlsxPageMargins)
		s.PageMargins = pm
	}

	for _, opt := range opts {
		opt.setPageMargins(pm)
	}
	return err
}

// GetPageMargins provides a function to get worksheet page margins.
//
// Available options:
//   PageMarginBotom(float64)
//   PageMarginFooter(float64)
//   PageMarginHeader(float64)
//   PageMarginLeft(float64)
//   PageMarginRight(float64)
//   PageMarginTop(float64)
func (f *File) GetPageMargins(sheet string, opts ...PageMarginsOptionsPtr) error {
	s, err := f.workSheetReader(sheet)
	if err != nil {
		return err
	}
	pm := s.PageMargins

	for _, opt := range opts {
		opt.getPageMargins(pm)
	}
	return err
}
