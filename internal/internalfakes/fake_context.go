// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"image"
	"image/color"
	"sync"

	"github.com/fogleman/gg"
	"github.com/petewall/eink-radiator-image-source-text/internal"
)

type FakeContext struct {
	DrawStringStub        func(string, float64, float64)
	drawStringMutex       sync.RWMutex
	drawStringArgsForCall []struct {
		arg1 string
		arg2 float64
		arg3 float64
	}
	DrawStringAnchoredStub        func(string, float64, float64, float64, float64)
	drawStringAnchoredMutex       sync.RWMutex
	drawStringAnchoredArgsForCall []struct {
		arg1 string
		arg2 float64
		arg3 float64
		arg4 float64
		arg5 float64
	}
	DrawStringWrappedStub        func(string, float64, float64, float64, float64, float64, float64, gg.Align)
	drawStringWrappedMutex       sync.RWMutex
	drawStringWrappedArgsForCall []struct {
		arg1 string
		arg2 float64
		arg3 float64
		arg4 float64
		arg5 float64
		arg6 float64
		arg7 float64
		arg8 gg.Align
	}
	ImageStub        func() image.Image
	imageMutex       sync.RWMutex
	imageArgsForCall []struct {
	}
	imageReturns struct {
		result1 image.Image
	}
	imageReturnsOnCall map[int]struct {
		result1 image.Image
	}
	LoadFontFaceStub        func(string, float64) error
	loadFontFaceMutex       sync.RWMutex
	loadFontFaceArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	loadFontFaceReturns struct {
		result1 error
	}
	loadFontFaceReturnsOnCall map[int]struct {
		result1 error
	}
	MeasureMultilineStringStub        func(string, float64) (float64, float64)
	measureMultilineStringMutex       sync.RWMutex
	measureMultilineStringArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	measureMultilineStringReturns struct {
		result1 float64
		result2 float64
	}
	measureMultilineStringReturnsOnCall map[int]struct {
		result1 float64
		result2 float64
	}
	MeasureStringStub        func(string) (float64, float64)
	measureStringMutex       sync.RWMutex
	measureStringArgsForCall []struct {
		arg1 string
	}
	measureStringReturns struct {
		result1 float64
		result2 float64
	}
	measureStringReturnsOnCall map[int]struct {
		result1 float64
		result2 float64
	}
	SetColorStub        func(color.Color)
	setColorMutex       sync.RWMutex
	setColorArgsForCall []struct {
		arg1 color.Color
	}
	WordWrapStub        func(string, float64) []string
	wordWrapMutex       sync.RWMutex
	wordWrapArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	wordWrapReturns struct {
		result1 []string
	}
	wordWrapReturnsOnCall map[int]struct {
		result1 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContext) DrawString(arg1 string, arg2 float64, arg3 float64) {
	fake.drawStringMutex.Lock()
	fake.drawStringArgsForCall = append(fake.drawStringArgsForCall, struct {
		arg1 string
		arg2 float64
		arg3 float64
	}{arg1, arg2, arg3})
	stub := fake.DrawStringStub
	fake.recordInvocation("DrawString", []interface{}{arg1, arg2, arg3})
	fake.drawStringMutex.Unlock()
	if stub != nil {
		fake.DrawStringStub(arg1, arg2, arg3)
	}
}

func (fake *FakeContext) DrawStringCallCount() int {
	fake.drawStringMutex.RLock()
	defer fake.drawStringMutex.RUnlock()
	return len(fake.drawStringArgsForCall)
}

func (fake *FakeContext) DrawStringCalls(stub func(string, float64, float64)) {
	fake.drawStringMutex.Lock()
	defer fake.drawStringMutex.Unlock()
	fake.DrawStringStub = stub
}

func (fake *FakeContext) DrawStringArgsForCall(i int) (string, float64, float64) {
	fake.drawStringMutex.RLock()
	defer fake.drawStringMutex.RUnlock()
	argsForCall := fake.drawStringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeContext) DrawStringAnchored(arg1 string, arg2 float64, arg3 float64, arg4 float64, arg5 float64) {
	fake.drawStringAnchoredMutex.Lock()
	fake.drawStringAnchoredArgsForCall = append(fake.drawStringAnchoredArgsForCall, struct {
		arg1 string
		arg2 float64
		arg3 float64
		arg4 float64
		arg5 float64
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.DrawStringAnchoredStub
	fake.recordInvocation("DrawStringAnchored", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.drawStringAnchoredMutex.Unlock()
	if stub != nil {
		fake.DrawStringAnchoredStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakeContext) DrawStringAnchoredCallCount() int {
	fake.drawStringAnchoredMutex.RLock()
	defer fake.drawStringAnchoredMutex.RUnlock()
	return len(fake.drawStringAnchoredArgsForCall)
}

func (fake *FakeContext) DrawStringAnchoredCalls(stub func(string, float64, float64, float64, float64)) {
	fake.drawStringAnchoredMutex.Lock()
	defer fake.drawStringAnchoredMutex.Unlock()
	fake.DrawStringAnchoredStub = stub
}

func (fake *FakeContext) DrawStringAnchoredArgsForCall(i int) (string, float64, float64, float64, float64) {
	fake.drawStringAnchoredMutex.RLock()
	defer fake.drawStringAnchoredMutex.RUnlock()
	argsForCall := fake.drawStringAnchoredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeContext) DrawStringWrapped(arg1 string, arg2 float64, arg3 float64, arg4 float64, arg5 float64, arg6 float64, arg7 float64, arg8 gg.Align) {
	fake.drawStringWrappedMutex.Lock()
	fake.drawStringWrappedArgsForCall = append(fake.drawStringWrappedArgsForCall, struct {
		arg1 string
		arg2 float64
		arg3 float64
		arg4 float64
		arg5 float64
		arg6 float64
		arg7 float64
		arg8 gg.Align
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	stub := fake.DrawStringWrappedStub
	fake.recordInvocation("DrawStringWrapped", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.drawStringWrappedMutex.Unlock()
	if stub != nil {
		fake.DrawStringWrappedStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
}

func (fake *FakeContext) DrawStringWrappedCallCount() int {
	fake.drawStringWrappedMutex.RLock()
	defer fake.drawStringWrappedMutex.RUnlock()
	return len(fake.drawStringWrappedArgsForCall)
}

func (fake *FakeContext) DrawStringWrappedCalls(stub func(string, float64, float64, float64, float64, float64, float64, gg.Align)) {
	fake.drawStringWrappedMutex.Lock()
	defer fake.drawStringWrappedMutex.Unlock()
	fake.DrawStringWrappedStub = stub
}

func (fake *FakeContext) DrawStringWrappedArgsForCall(i int) (string, float64, float64, float64, float64, float64, float64, gg.Align) {
	fake.drawStringWrappedMutex.RLock()
	defer fake.drawStringWrappedMutex.RUnlock()
	argsForCall := fake.drawStringWrappedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8
}

func (fake *FakeContext) Image() image.Image {
	fake.imageMutex.Lock()
	ret, specificReturn := fake.imageReturnsOnCall[len(fake.imageArgsForCall)]
	fake.imageArgsForCall = append(fake.imageArgsForCall, struct {
	}{})
	stub := fake.ImageStub
	fakeReturns := fake.imageReturns
	fake.recordInvocation("Image", []interface{}{})
	fake.imageMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContext) ImageCallCount() int {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return len(fake.imageArgsForCall)
}

func (fake *FakeContext) ImageCalls(stub func() image.Image) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = stub
}

func (fake *FakeContext) ImageReturns(result1 image.Image) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	fake.imageReturns = struct {
		result1 image.Image
	}{result1}
}

func (fake *FakeContext) ImageReturnsOnCall(i int, result1 image.Image) {
	fake.imageMutex.Lock()
	defer fake.imageMutex.Unlock()
	fake.ImageStub = nil
	if fake.imageReturnsOnCall == nil {
		fake.imageReturnsOnCall = make(map[int]struct {
			result1 image.Image
		})
	}
	fake.imageReturnsOnCall[i] = struct {
		result1 image.Image
	}{result1}
}

func (fake *FakeContext) LoadFontFace(arg1 string, arg2 float64) error {
	fake.loadFontFaceMutex.Lock()
	ret, specificReturn := fake.loadFontFaceReturnsOnCall[len(fake.loadFontFaceArgsForCall)]
	fake.loadFontFaceArgsForCall = append(fake.loadFontFaceArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	stub := fake.LoadFontFaceStub
	fakeReturns := fake.loadFontFaceReturns
	fake.recordInvocation("LoadFontFace", []interface{}{arg1, arg2})
	fake.loadFontFaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContext) LoadFontFaceCallCount() int {
	fake.loadFontFaceMutex.RLock()
	defer fake.loadFontFaceMutex.RUnlock()
	return len(fake.loadFontFaceArgsForCall)
}

func (fake *FakeContext) LoadFontFaceCalls(stub func(string, float64) error) {
	fake.loadFontFaceMutex.Lock()
	defer fake.loadFontFaceMutex.Unlock()
	fake.LoadFontFaceStub = stub
}

func (fake *FakeContext) LoadFontFaceArgsForCall(i int) (string, float64) {
	fake.loadFontFaceMutex.RLock()
	defer fake.loadFontFaceMutex.RUnlock()
	argsForCall := fake.loadFontFaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContext) LoadFontFaceReturns(result1 error) {
	fake.loadFontFaceMutex.Lock()
	defer fake.loadFontFaceMutex.Unlock()
	fake.LoadFontFaceStub = nil
	fake.loadFontFaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContext) LoadFontFaceReturnsOnCall(i int, result1 error) {
	fake.loadFontFaceMutex.Lock()
	defer fake.loadFontFaceMutex.Unlock()
	fake.LoadFontFaceStub = nil
	if fake.loadFontFaceReturnsOnCall == nil {
		fake.loadFontFaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadFontFaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContext) MeasureMultilineString(arg1 string, arg2 float64) (float64, float64) {
	fake.measureMultilineStringMutex.Lock()
	ret, specificReturn := fake.measureMultilineStringReturnsOnCall[len(fake.measureMultilineStringArgsForCall)]
	fake.measureMultilineStringArgsForCall = append(fake.measureMultilineStringArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	stub := fake.MeasureMultilineStringStub
	fakeReturns := fake.measureMultilineStringReturns
	fake.recordInvocation("MeasureMultilineString", []interface{}{arg1, arg2})
	fake.measureMultilineStringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContext) MeasureMultilineStringCallCount() int {
	fake.measureMultilineStringMutex.RLock()
	defer fake.measureMultilineStringMutex.RUnlock()
	return len(fake.measureMultilineStringArgsForCall)
}

func (fake *FakeContext) MeasureMultilineStringCalls(stub func(string, float64) (float64, float64)) {
	fake.measureMultilineStringMutex.Lock()
	defer fake.measureMultilineStringMutex.Unlock()
	fake.MeasureMultilineStringStub = stub
}

func (fake *FakeContext) MeasureMultilineStringArgsForCall(i int) (string, float64) {
	fake.measureMultilineStringMutex.RLock()
	defer fake.measureMultilineStringMutex.RUnlock()
	argsForCall := fake.measureMultilineStringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContext) MeasureMultilineStringReturns(result1 float64, result2 float64) {
	fake.measureMultilineStringMutex.Lock()
	defer fake.measureMultilineStringMutex.Unlock()
	fake.MeasureMultilineStringStub = nil
	fake.measureMultilineStringReturns = struct {
		result1 float64
		result2 float64
	}{result1, result2}
}

func (fake *FakeContext) MeasureMultilineStringReturnsOnCall(i int, result1 float64, result2 float64) {
	fake.measureMultilineStringMutex.Lock()
	defer fake.measureMultilineStringMutex.Unlock()
	fake.MeasureMultilineStringStub = nil
	if fake.measureMultilineStringReturnsOnCall == nil {
		fake.measureMultilineStringReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 float64
		})
	}
	fake.measureMultilineStringReturnsOnCall[i] = struct {
		result1 float64
		result2 float64
	}{result1, result2}
}

func (fake *FakeContext) MeasureString(arg1 string) (float64, float64) {
	fake.measureStringMutex.Lock()
	ret, specificReturn := fake.measureStringReturnsOnCall[len(fake.measureStringArgsForCall)]
	fake.measureStringArgsForCall = append(fake.measureStringArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MeasureStringStub
	fakeReturns := fake.measureStringReturns
	fake.recordInvocation("MeasureString", []interface{}{arg1})
	fake.measureStringMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContext) MeasureStringCallCount() int {
	fake.measureStringMutex.RLock()
	defer fake.measureStringMutex.RUnlock()
	return len(fake.measureStringArgsForCall)
}

func (fake *FakeContext) MeasureStringCalls(stub func(string) (float64, float64)) {
	fake.measureStringMutex.Lock()
	defer fake.measureStringMutex.Unlock()
	fake.MeasureStringStub = stub
}

func (fake *FakeContext) MeasureStringArgsForCall(i int) string {
	fake.measureStringMutex.RLock()
	defer fake.measureStringMutex.RUnlock()
	argsForCall := fake.measureStringArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContext) MeasureStringReturns(result1 float64, result2 float64) {
	fake.measureStringMutex.Lock()
	defer fake.measureStringMutex.Unlock()
	fake.MeasureStringStub = nil
	fake.measureStringReturns = struct {
		result1 float64
		result2 float64
	}{result1, result2}
}

func (fake *FakeContext) MeasureStringReturnsOnCall(i int, result1 float64, result2 float64) {
	fake.measureStringMutex.Lock()
	defer fake.measureStringMutex.Unlock()
	fake.MeasureStringStub = nil
	if fake.measureStringReturnsOnCall == nil {
		fake.measureStringReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 float64
		})
	}
	fake.measureStringReturnsOnCall[i] = struct {
		result1 float64
		result2 float64
	}{result1, result2}
}

func (fake *FakeContext) SetColor(arg1 color.Color) {
	fake.setColorMutex.Lock()
	fake.setColorArgsForCall = append(fake.setColorArgsForCall, struct {
		arg1 color.Color
	}{arg1})
	stub := fake.SetColorStub
	fake.recordInvocation("SetColor", []interface{}{arg1})
	fake.setColorMutex.Unlock()
	if stub != nil {
		fake.SetColorStub(arg1)
	}
}

func (fake *FakeContext) SetColorCallCount() int {
	fake.setColorMutex.RLock()
	defer fake.setColorMutex.RUnlock()
	return len(fake.setColorArgsForCall)
}

func (fake *FakeContext) SetColorCalls(stub func(color.Color)) {
	fake.setColorMutex.Lock()
	defer fake.setColorMutex.Unlock()
	fake.SetColorStub = stub
}

func (fake *FakeContext) SetColorArgsForCall(i int) color.Color {
	fake.setColorMutex.RLock()
	defer fake.setColorMutex.RUnlock()
	argsForCall := fake.setColorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContext) WordWrap(arg1 string, arg2 float64) []string {
	fake.wordWrapMutex.Lock()
	ret, specificReturn := fake.wordWrapReturnsOnCall[len(fake.wordWrapArgsForCall)]
	fake.wordWrapArgsForCall = append(fake.wordWrapArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	stub := fake.WordWrapStub
	fakeReturns := fake.wordWrapReturns
	fake.recordInvocation("WordWrap", []interface{}{arg1, arg2})
	fake.wordWrapMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeContext) WordWrapCallCount() int {
	fake.wordWrapMutex.RLock()
	defer fake.wordWrapMutex.RUnlock()
	return len(fake.wordWrapArgsForCall)
}

func (fake *FakeContext) WordWrapCalls(stub func(string, float64) []string) {
	fake.wordWrapMutex.Lock()
	defer fake.wordWrapMutex.Unlock()
	fake.WordWrapStub = stub
}

func (fake *FakeContext) WordWrapArgsForCall(i int) (string, float64) {
	fake.wordWrapMutex.RLock()
	defer fake.wordWrapMutex.RUnlock()
	argsForCall := fake.wordWrapArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContext) WordWrapReturns(result1 []string) {
	fake.wordWrapMutex.Lock()
	defer fake.wordWrapMutex.Unlock()
	fake.WordWrapStub = nil
	fake.wordWrapReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeContext) WordWrapReturnsOnCall(i int, result1 []string) {
	fake.wordWrapMutex.Lock()
	defer fake.wordWrapMutex.Unlock()
	fake.WordWrapStub = nil
	if fake.wordWrapReturnsOnCall == nil {
		fake.wordWrapReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.wordWrapReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeContext) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.drawStringMutex.RLock()
	defer fake.drawStringMutex.RUnlock()
	fake.drawStringAnchoredMutex.RLock()
	defer fake.drawStringAnchoredMutex.RUnlock()
	fake.drawStringWrappedMutex.RLock()
	defer fake.drawStringWrappedMutex.RUnlock()
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	fake.loadFontFaceMutex.RLock()
	defer fake.loadFontFaceMutex.RUnlock()
	fake.measureMultilineStringMutex.RLock()
	defer fake.measureMultilineStringMutex.RUnlock()
	fake.measureStringMutex.RLock()
	defer fake.measureStringMutex.RUnlock()
	fake.setColorMutex.RLock()
	defer fake.setColorMutex.RUnlock()
	fake.wordWrapMutex.RLock()
	defer fake.wordWrapMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContext) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ internal.Context = new(FakeContext)
