package main

import (
  capn "github.com/glycerine/go-capnproto"
  "io"
)




func (s *Clue) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	ClueGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *Clue) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootClueCapn(capMsg)
      ClueCapnToGo(z, s)
   return nil
}



func ClueCapnToGo(src ClueCapn, dest *Clue) *Clue {
  if dest == nil {
    dest = &Clue{}
  }
  dest.Term = src.Term()
  dest.Intro = src.Intro()
  dest.Potency = src.Potency()

  return dest
}



func ClueGoToCapn(seg *capn.Segment, src *Clue) ClueCapn {
  dest := AutoNewClueCapn(seg)
  dest.SetTerm(src.Term)
  dest.SetIntro(src.Intro)
  dest.SetPotency(src.Potency)

  return dest
}



func (s *Rev) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	RevGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *Rev) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootRevCapn(capMsg)
      RevCapnToGo(z, s)
   return nil
}



func RevCapnToGo(src RevCapn, dest *Rev) *Rev {
  if dest == nil {
    dest = &Rev{}
  }
  dest.DocId = src.DocId()
  dest.Rank = src.Rank()
  dest.InMeta = src.InMeta()

  return dest
}



func RevGoToCapn(seg *capn.Segment, src *Rev) RevCapn {
  dest := AutoNewRevCapn(seg)
  dest.SetDocId(src.DocId)
  dest.SetRank(src.Rank)
  dest.SetInMeta(src.InMeta)

  return dest
}



func (s *Revs) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	RevsGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *Revs) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootRevsCapn(capMsg)
      RevsCapnToGo(z, s)
   return nil
}



func RevsCapnToGo(src RevsCapn, dest *Revs) *Revs {
  if dest == nil {
    dest = &Revs{}
  }

  var n int

    // Data
	n = src.Data().Len()
	dest.Data = make([]Rev, n)
	for i := 0; i < n; i++ {
        dest.Data[i] = *RevCapnToGo(src.Data().At(i), nil)
    }


  return dest
}



func RevsGoToCapn(seg *capn.Segment, src *Revs) RevsCapn {
  dest := AutoNewRevsCapn(seg)

  // Data -> RevCapn (go slice to capn list)
  if len(src.Data) > 0 {
		typedList := NewRevCapnList(seg, len(src.Data))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Data {
			plist.Set(i, capn.Object(RevGoToCapn(seg, &ele)))
			i++
		}
		dest.SetData(typedList)
	}

  return dest
}



func (s *Shotgun) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	ShotgunGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *Shotgun) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootShotgunCapn(capMsg)
      ShotgunCapnToGo(z, s)
   return nil
}



func ShotgunCapnToGo(src ShotgunCapn, dest *Shotgun) *Shotgun {
  if dest == nil {
    dest = &Shotgun{}
  }
  dest.Term = src.Term()
  dest.Potency = src.Potency()

  return dest
}



func ShotgunGoToCapn(seg *capn.Segment, src *Shotgun) ShotgunCapn {
  dest := AutoNewShotgunCapn(seg)
  dest.SetTerm(src.Term)
  dest.SetPotency(src.Potency)

  return dest
}



func (s *Term) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	TermGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 


func (s *Term) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
  	}
  	z := ReadRootTermCapn(capMsg)
      TermCapnToGo(z, s)
   return nil
}



func TermCapnToGo(src TermCapn, dest *Term) *Term {
  if dest == nil {
    dest = &Term{}
  }
  dest.TermStr = src.TermStr()
  dest.Slot = src.Slot()
  dest.NumDocuments = src.NumDocuments()
  dest.NumWords = src.NumWords()

  var n int

    // Shotgun
	n = src.Shotgun().Len()
	dest.Shotgun = make([]Shotgun, n)
	for i := 0; i < n; i++ {
        dest.Shotgun[i] = *ShotgunCapnToGo(src.Shotgun().At(i), nil)
    }


    // Clues
	n = src.Clues().Len()
	dest.Clues = make([]Clue, n)
	for i := 0; i < n; i++ {
        dest.Clues[i] = *ClueCapnToGo(src.Clues().At(i), nil)
    }

  dest.InteractionsPos = src.InteractionsPos()
  dest.InteractionsNeg = src.InteractionsNeg()
  dest.HardcodedScore = src.HardcodedScore()
  dest.Infogain = src.Infogain()

  return dest
}



func TermGoToCapn(seg *capn.Segment, src *Term) TermCapn {
  dest := AutoNewTermCapn(seg)
  dest.SetTermStr(src.TermStr)
  dest.SetSlot(src.Slot)
  dest.SetNumDocuments(src.NumDocuments)
  dest.SetNumWords(src.NumWords)

  // Shotgun -> ShotgunCapn (go slice to capn list)
  if len(src.Shotgun) > 0 {
		typedList := NewShotgunCapnList(seg, len(src.Shotgun))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Shotgun {
			plist.Set(i, capn.Object(ShotgunGoToCapn(seg, &ele)))
			i++
		}
		dest.SetShotgun(typedList)
	}

  // Clues -> ClueCapn (go slice to capn list)
  if len(src.Clues) > 0 {
		typedList := NewClueCapnList(seg, len(src.Clues))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Clues {
			plist.Set(i, capn.Object(ClueGoToCapn(seg, &ele)))
			i++
		}
		dest.SetClues(typedList)
	}
  dest.SetInteractionsPos(src.InteractionsPos)
  dest.SetInteractionsNeg(src.InteractionsNeg)
  dest.SetHardcodedScore(src.HardcodedScore)
  dest.SetInfogain(src.Infogain)

  return dest
}



func SliceClueToClueCapnList(seg *capn.Segment, m []Clue) ClueCapn_List {
	lst := NewClueCapnList(seg, len(m))
	for i := range m {
		lst.Set(i, ClueGoToCapn(seg, &m[i]))
	}
	return lst
}



func ClueCapnListToSliceClue(p ClueCapn_List) []Clue {
	v := make([]Clue, p.Len())
	for i := range v {
        ClueCapnToGo(p.At(i), &v[i])
	}
	return v
}



func SliceRevToRevCapnList(seg *capn.Segment, m []Rev) RevCapn_List {
	lst := NewRevCapnList(seg, len(m))
	for i := range m {
		lst.Set(i, RevGoToCapn(seg, &m[i]))
	}
	return lst
}



func RevCapnListToSliceRev(p RevCapn_List) []Rev {
	v := make([]Rev, p.Len())
	for i := range v {
        RevCapnToGo(p.At(i), &v[i])
	}
	return v
}



func SliceShotgunToShotgunCapnList(seg *capn.Segment, m []Shotgun) ShotgunCapn_List {
	lst := NewShotgunCapnList(seg, len(m))
	for i := range m {
		lst.Set(i, ShotgunGoToCapn(seg, &m[i]))
	}
	return lst
}



func ShotgunCapnListToSliceShotgun(p ShotgunCapn_List) []Shotgun {
	v := make([]Shotgun, p.Len())
	for i := range v {
        ShotgunCapnToGo(p.At(i), &v[i])
	}
	return v
}
